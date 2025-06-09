package oauth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

func authorizeHandler(c *gin.Context) {
	var location string
	var responseType string
	var clientID string
	var redirectURI string
	var scope string
	var state string

	session := sessions.Default(c)
	userPublicID := session.Get("userPublicID")
	nextPath := url.QueryEscape(fmt.Sprintf("%s?%s", c.Request.URL.Path, c.Request.URL.RawQuery))
	if userPublicID == nil {
		location = fmt.Sprintf("/signin?_=%s", nextPath)
		c.Redirect(http.StatusFound, location)
		return
	}
	user := services.FindUserByPublicID(userPublicID.(string))
	if user.ID == 0 {
		session.Delete("userPublicID")
		session.Save()
		location = fmt.Sprintf("/signin?_=%s", nextPath)
		c.Redirect(http.StatusFound, location)
		return
	}

	responseType = c.Query("response_type")
	clientID = c.Query("client_id")
	redirectURI = c.Query("redirect_uri")
	scope = c.Query("scope")
	state = c.Query("state")

	if redirectURI == "" {
		redirectURI = "/error"
	}

	client := services.FindClientByKey(clientID)
	if client.ID == 0 {
		// REFACTOR This scenario is the trickiest one
		// redirectURI = "/error"
		// location = fmt.Sprintf("%s?error=%s&state=%s", redirectURI, UnauthorizedClient, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": nil,
			"ErrorCode": UnauthorizedClient,
		})
		return
	}

	if scope != models.PublicScope && scope != models.ReadScope && scope != models.WriteScope && scope != models.OpenIDScope {
		scope = models.PublicScope
	}

	switch responseType {
	// Authorization Code Grant
	case Code:
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
				// In this scenario, the user requested to deny access; it's not the client application's fault
				// The client application is safe, so the user may proceed (client application must handle this)
				location = fmt.Sprintf(errorURI, redirectURI, AccessDenied, state)
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
			})
			if err != nil {
				location = fmt.Sprintf(errorURI, redirectURI, result["error"], result["state"])
				// Previous return: c.HTML(http.StatusFound, location)
				c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
					"Title":     " - Authorization Error",
					"Internal":  true,
					"ProceedTo": location,
					"ErrorCode": result["error"],
				})
			} else {
				location = fmt.Sprintf("%s?code=%s&scope=%s&state=%s",
					redirectURI, result["code"], result["scope"], result["state"])
				c.Redirect(http.StatusFound, location)
			}
		} else {
			c.String(http.StatusNotFound, "404 Not Found")
		}
	// Implicit Grant
	case Token:
		location = fmt.Sprintf(errorURI, redirectURI, UnsupportedResponseType, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": location,
			"ErrorCode": UnsupportedResponseType,
		})
	default:
		location = fmt.Sprintf(errorURI, redirectURI, InvalidRequest, state)
		// Previous return: c.HTML(http.StatusFound, location)
		c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
			"Title":     " - Authorization Error",
			"Internal":  true,
			"ProceedTo": location,
			"ErrorCode": InvalidRequest,
		})
	}
}
