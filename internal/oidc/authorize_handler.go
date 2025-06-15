package oidc

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/services"
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
	userPublicID := session.Get("user_public_id")
	nextPath := url.QueryEscape(fmt.Sprintf("%s?%s", c.Request.URL.Path, c.Request.URL.RawQuery))
	if userPublicID == nil {
		location = fmt.Sprintf("/signin?_=%s", nextPath)
		c.Redirect(http.StatusFound, location)
		return
	}
	user := services.FindUserByPublicID(userPublicID.(string))
	if user.IsNewRecord() {
		session.Delete("user_public_id")
		session.Save()
		location = fmt.Sprintf("/signin?_=%s", nextPath)
		c.Redirect(http.StatusFound, location)
		return
	}

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

	client := services.FindClientByKey(clientKey)
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
			result, err := ImplicitFlowIDToken(utils.H{
				"response_type": responseType,
				"client":        client,
				"user":          user,
				"ip":            c.Request.RemoteAddr,
				"userAgent":     c.Request.UserAgent(),
				"redirect_uri":  redirectURI,
				"scope":         scope,
				"state":         state,
				"nonce":         nonce,
				"response_mode": responseMode,
				"issuer":        shared.GetBaseUrl(c),
			})
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
		activeSessions := services.ActiveSessionsForClient(client.ID, user.ID)
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
			result, err := AuthorizationCodeGrant(utils.H{
				"response_type": responseType,
				"client":        client,
				"user":          user,
				"ip":            c.Request.RemoteAddr,
				"userAgent":     c.Request.UserAgent(),
				"redirect_uri":  redirectURI,
				"scope":         scope,
				"state":         state,
				"nonce":         nonce,
				"response_mode": responseMode,
			})
			processResponseForAuthorizeHandlerCode(c, result, err)
			return
		} else {
			c.String(http.StatusNotFound, "404 Not Found")
		}
	// Implicit Grant (OAuth)
	// Hybrid Flow (OIDC)
	case shared.Token, shared.CodeIDToken:
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
