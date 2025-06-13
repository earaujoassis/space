package oidc

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
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
	if user.ID == 0 {
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

	if responseType == "" || clientKey == "" || redirectURI == "" || scope == "" || nonce == "" {
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": nil,
			"ErrorCode": shared.InvalidRequest,
		})
	}

	client := services.FindClientByKey(clientKey)
	if client.ID == 0 {
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

	if scope == "" || !models.HasValidScopes(strings.Split(scope, " ")) || !strings.Contains(scope, models.OpenIDScope) {
		location = fmt.Sprintf(shared.ErrorURI, redirectURI, shared.InvalidScope, state)
		c.Redirect(http.StatusFound, location)
		return
	}

	switch responseType {
	// Implicit Flow (OIDC)
	case shared.IdToken:
		response_mode := c.Query("response_mode")
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
				location = fmt.Sprintf(shared.ErrorURI, redirectURI, shared.AccessDenied, state)
				c.Redirect(http.StatusFound, location)
				return
			}
			result, err := ImplicitFlowIdToken(utils.H{
				"response_type": responseType,
				"client":        client,
				"user":          user,
				"ip":            c.Request.RemoteAddr,
				"userAgent":     c.Request.UserAgent(),
				"redirect_uri":  redirectURI,
				"scope":         scope,
				"state":         state,
				"nonce":         nonce,
				"response_mode": response_mode,
				"issuer":        shared.GetBaseUrl(c),
			})
			if err != nil {
				location = fmt.Sprintf(shared.ErrorURI, redirectURI, result["error"], result["state"])
				// Previous return: c.HTML(http.StatusFound, location)
				c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
					"Title":     " - Authorization Error",
					"Internal":  true,
					"ProceedTo": location,
					"ErrorCode": result["error"],
				})
			} else {
				if response_mode == "form_post" || response_mode == "" {
					c.HTML(http.StatusOK, "form_post", utils.H{
						"Callback": redirectURI,
						"IdToken":  result["id_token"],
						"State":  result["state"],
					})
				}
			}
		} else {
			c.String(http.StatusNotFound, "404 Not Found")
		}
	// Authorization Code Grant (OAuth)
	// Implicit Grant (OAuth)
	// Hybrid Flow (OIDC)
	case shared.Code, shared.Token, shared.CodeIdToken:
		location = fmt.Sprintf(shared.ErrorURI, redirectURI, shared.UnsupportedResponseType, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": location,
			"ErrorCode": shared.UnsupportedResponseType,
		})
	default:
		location = fmt.Sprintf(shared.ErrorURI, redirectURI, shared.InvalidRequest, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": location,
			"ErrorCode": shared.InvalidRequest,
		})
	}
}
