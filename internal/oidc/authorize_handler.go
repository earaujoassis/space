package oidc

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func authorizeHandler(c *gin.Context) {
	var location string
	var responseType string
	var clientKey string
	var redirectURI string
	var scope string
	var state string
	var nonce string

	session := sessions.Default(c)
	applicationTokenInterface := session.Get(shared.CookieSessionKey)
	applicationToken := utils.StringValue(applicationTokenInterface)
	nextPath := url.QueryEscape(fmt.Sprintf("%s?%s", c.Request.URL.Path, c.Request.URL.RawQuery))
	location = fmt.Sprintf("/signin?_=%s", nextPath)
	if applicationTokenInterface == nil {
		c.Redirect(http.StatusFound, location)
		return
	}
	if !security.ValidToken(applicationToken) {
		session.Delete(shared.CookieSessionKey)
		session.Save()
		c.Redirect(http.StatusFound, location)
		return
	}
	repositories := ioc.GetRepositories(c)
	applicationSession := repositories.Sessions().FindByToken(applicationToken, models.ApplicationToken)
	if applicationSession.IsNewRecord() || applicationSession.User.IsNewRecord() {
		session.Delete(shared.CookieSessionKey)
		session.Save()
		c.Redirect(http.StatusFound, location)
		return
	}
	user := applicationSession.User

	responseType = c.Query("response_type")
	clientKey = c.Query("client_id")
	redirectURI = c.Query("redirect_uri")
	scope = c.Query("scope")
	state = c.Query("state")
	nonce = c.Query("nonce")

	if responseType == "" || clientKey == "" || redirectURI == "" || scope == "" {
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": nil,
			"ErrorCode": shared.InvalidRequest,
		})
	}

	client := repositories.Clients().FindByKey(clientKey)
	if client.IsNewRecord() {
		// WARNING This scenario is the trickiest one
		// It is not safe to return to the caller or redirect to callback
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": nil,
			"ErrorCode": shared.UnauthorizedClient,
		})
		return
	}

	if !slices.Contains(client.RedirectURI, redirectURI) {
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": nil,
			"ErrorCode": shared.InvalidRequest,
		})
		return
	}

	switch responseType {
	// Implicit Flow (OIDC)
	case shared.IDToken:
		if err := validateScopeForIDToken(c); err != nil {
			return
		}
		if err := validateResponseModeForIDToken(c); err != nil {
			return
		}
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "satellite", utils.H{
				"Title":     " - Authorize",
				"Satellite": "callisto",
				"Internal":  true,
				"Data": utils.H{
					"first_name":      user.FirstName,
					"last_name":       user.LastName,
					"client_name":     client.Name,
					"client_uri":      client.DefaultCanonicalURI(),
					"requested_scope": scope,
				},
			})
			return
		} else if c.Request.Method == "POST" {
			if c.PostForm("access_denied") == "true" {
				// The user requested to deny access
				processResponseForIDTokenAccessDenied(c)
				return
			}
			responseMode := c.Query("response_mode")
			result, err := ImplicitFlowIDToken(ImplicitFlowIDTokenParams{
				ResponseType: responseType,
				Client:       client,
				User:         user,
				IP:           c.ClientIP(),
				UserAgent:    c.Request.UserAgent(),
				RedirectURI:  redirectURI,
				Scope:        scope,
				State:        state,
				Nonce:        nonce,
				ResponseMode: responseMode,
				Issuer:       shared.GetBaseUrl(c),
			}, repositories)
			processResponseForAuthorizeHandlerIDToken(c, result, err)
			return
		} else {
			c.String(http.StatusNotFound, "404 Not Found")
		}
	// Authorization Code Grant (OIDC+OAuth)
	case shared.Code:
		if err := validateScopeForCode(c); err != nil {
			return
		}
		if err := validateResponseModeForCode(c); err != nil {
			return
		}
		activeSessions := repositories.Sessions().ActiveForClient(client, user)
		if c.Request.Method == "GET" && activeSessions == 0 {
			c.HTML(http.StatusOK, "satellite", utils.H{
				"Title":     " - Authorize",
				"Satellite": "callisto",
				"Internal":  true,
				"Data": utils.H{
					"first_name":      user.FirstName,
					"last_name":       user.LastName,
					"client_name":     client.Name,
					"client_uri":      client.DefaultCanonicalURI(),
					"requested_scope": scope,
				},
			})
			return
		} else if c.Request.Method == "POST" || (activeSessions > 0 && c.Request.Method == "GET") {
			if c.PostForm("access_denied") == "true" {
				// The user requested to deny access
				processResponseForCodeAccessDenied(c)
				return
			}
			responseMode := c.Query("response_mode")
			result, err := AuthorizationCodeGrant(AuthorizationCodeParams{
				ResponseType: responseType,
				Client:       client,
				User:         user,
				IP:           c.ClientIP(),
				UserAgent:    c.Request.UserAgent(),
				RedirectURI:  redirectURI,
				Scope:        scope,
				State:        state,
				Nonce:        nonce,
				ResponseMode: responseMode,
			}, repositories)
			if err == nil {
				notifier := ioc.GetNotifier(c)
				go notifier.Announce("user.authorization_granted", utils.H{
					"Email":      shared.GetUserDefaultEmailForNotifications(c, user),
					"FirstName":  user.FirstName,
					"ClientName": client.Name,
					"CreatedAt":  time.Now().UTC().Format(time.RFC850),
				})
			}
			processResponseForAuthorizeHandlerCode(c, result, err)
			return
		}
	// Implicit Grant (OAuth)
	// Hybrid Flow (OIDC)
	case shared.Token, shared.CodeIDToken, shared.IDTokenToken, shared.CodeToken, shared.CodeIDTokenToken:
		location = fmt.Sprintf(shared.ErrorQueryURI, redirectURI, shared.UnsupportedResponseType, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": location,
			"ErrorCode": shared.UnsupportedResponseType,
		})
	default:
		location = fmt.Sprintf(shared.ErrorQueryURI, redirectURI, shared.InvalidRequest, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": location,
			"ErrorCode": shared.InvalidRequest,
		})
	}
}
