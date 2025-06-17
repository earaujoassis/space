package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
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

	session := sessions.Default(c)
	userPublicID := session.Get("user_public_id")
	nextPath := url.QueryEscape(fmt.Sprintf("%s?%s", c.Request.URL.Path, c.Request.URL.RawQuery))
	if userPublicID == nil {
		location = fmt.Sprintf("/signin?_=%s", nextPath)
		c.Redirect(http.StatusFound, location)
		return
	}
	repositories := ioc.GetRepositories(c)
	user := repositories.Users().FindByPublicID(userPublicID.(string))
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

	if responseType == "" || clientKey == "" || redirectURI == "" {
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

	if (scope != "" && !models.HasValidScopes(utils.Scopes(scope))) || client.HasScope(models.OpenIDScope) {
		location = fmt.Sprintf(shared.ErrorQueryURI, redirectURI, shared.InvalidScope, state)
		c.Redirect(http.StatusFound, location)
		return
	}

	switch responseType {
	// Authorization Code Grant
	case shared.Code:
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
				location = fmt.Sprintf(shared.ErrorQueryURI, redirectURI, shared.AccessDenied, state)
				c.Redirect(http.StatusFound, location)
				return
			}
			result, err := AuthorizationCodeGrant(utils.H{
				"response_type": responseType,
				"client":        client,
				"user":          user,
				"ip":            c.Request.RemoteAddr,
				"userAgent":     c.Request.UserAgent(),
				"redirect_uri":  redirectURI,
				"scope":         scope,
				"state":         state,
			}, repositories)
			if err != nil {
				location = fmt.Sprintf(shared.ErrorQueryURI, redirectURI, result["error"], result["state"])
				// Previous return: c.HTML(http.StatusFound, location)
				c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
					"Title":     " - Authorization Error",
					"Internal":  true,
					"ProceedTo": location,
					"ErrorCode": result["error"],
				})
			} else {
				location = fmt.Sprintf("%s?code=%s&state=%s",
					redirectURI, result["code"], result["state"])
				c.Redirect(http.StatusFound, location)
			}
		}
	// Implicit Grant
	case shared.Token:
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
