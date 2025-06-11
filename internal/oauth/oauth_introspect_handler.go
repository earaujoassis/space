package oauth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

func introspectHandler(c *gin.Context) {
	var token = c.PostForm("token")
	var tokenTypeHint = c.PostForm("token_type_hint")
	var session models.Session
	var introspectType string
	var client models.Client

	baseURL := getBaseUrl(c)
	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	if client = ClientAuthentication(authorizationBasic); client.ID == 0 {
		c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Cannot fulfill token request",
			"error":    AccessDenied,
		})
		return
	}

	if tokenTypeHint != models.AccessToken && tokenTypeHint != models.RefreshToken && tokenTypeHint != "" {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Missing or invalid token hint parameter",
			"error":    InvalidRequest,
		})
		return
	}

	if !security.ValidToken(token) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Session instropection failed",
			"error":    InvalidRequest,
		})
		return
	}

	switch tokenTypeHint {
	case models.AccessToken:
		session = services.FindSessionByToken(token, models.AccessToken)
		introspectType = models.AccessToken
	case models.RefreshToken:
		session = services.FindSessionByToken(token, models.RefreshToken)
		introspectType = models.RefreshToken
	default:
		session = services.FindSessionByToken(token, models.AccessToken)
		introspectType = models.AccessToken
		if session.ID == 0 {
			session = services.FindSessionByToken(token, models.RefreshToken)
			introspectType = models.RefreshToken
		}
	}

	if session.ID == 0 {
		session = services.FindSessionByToken(token, models.AccessToken)
		introspectType = models.AccessToken
		if session.ID == 0 {
			session = services.FindSessionByToken(token, models.RefreshToken)
			introspectType = models.RefreshToken
		}
	}

	if session.ID == 0 {
		c.JSON(http.StatusOK, utils.H{
			"active": false,
		})
		return
	}

	if session.Client.ID != client.ID {
		c.JSON(http.StatusOK, utils.H{
			"active": false,
		})
		return
	}

	user := session.User
	client = session.Client

	introspectionData := utils.H{
		"active": true,
		"scope": session.Scopes,
		"client_id": client.Key,
		"username": user.Username,
		"exp": session.ExpiresIn,
		"iat": session.Moment,
		// "nbf": (not before) not defined nor required
		"sub": user.PublicID,
		// "aud": (audience) not defined nor required
		"iss": baseURL,
		// "jti": (String identifier for the token) not defined nor required
		"user_id": user.PublicID,
		"roles": []string{ "user" },
		// "email": user.Email,
	}

	if introspectType == models.AccessToken {
		introspectionData["token_type"] = "Bearer"
	}

	c.JSON(http.StatusOK, introspectionData)
}
