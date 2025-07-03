package oauth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func introspectHandler(c *gin.Context) {
	var token = c.PostForm("token")
	var tokenTypeHint = c.PostForm("token_type_hint")
	var session models.Session
	var client models.Client
	var introspectType string

	baseURL := shared.GetBaseUrl(c)
	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	key, secret := shared.BasicAuthDecode(authorizationBasic)
	repositories := ioc.GetRepositories(c)
	if client = repositories.Clients().Authentication(key, secret); client.IsNewRecord() {
		c.Header("WWW-Authenticate", "Basic realm=\"OAuth\"")
		c.JSON(http.StatusUnauthorized, utils.H{
			"error":             shared.InvalidClient,
			"error_description": "Client authentication failed",
		})
		return
	}

	if token == "" {
		c.JSON(http.StatusBadRequest, utils.H{
			"error":             shared.InvalidClient,
			"error_description": "Missing token parameter",
		})
		return
	}

	if !security.ValidToken(token) {
		c.JSON(http.StatusOK, utils.H{
			"active": false,
		})
		return
	}

	if !security.ValidToken(token) {
		c.JSON(http.StatusOK, utils.H{
			"active": false,
		})
		return
	}

	switch tokenTypeHint {
	case models.AccessToken:
		session = repositories.Sessions().FindByToken(token, models.AccessToken)
		introspectType = models.AccessToken
	case models.RefreshToken:
		session = repositories.Sessions().FindByToken(token, models.RefreshToken)
		introspectType = models.RefreshToken
	}

	if session.IsNewRecord() {
		session = repositories.Sessions().FindByToken(token, models.AccessToken)
		introspectType = models.AccessToken
		if session.IsNewRecord() {
			session = repositories.Sessions().FindByToken(token, models.RefreshToken)
			introspectType = models.RefreshToken
		}
	}

	if session.IsNewRecord() || session.Client.ID != client.ID {
		c.JSON(http.StatusOK, utils.H{
			"active": false,
		})
		return
	}

	user := session.User
	client = session.Client
	introspectionData := utils.H{
		"active":    true,
		"scope":     session.Scopes,
		"client_id": client.Key,
		"username":  user.Username,
		"exp":       session.Moment + session.ExpiresIn,
		"iat":       session.Moment,
		// "nbf": (not before) not defined nor required
		"sub": user.PublicID,
		// "aud": (audience) not defined nor required
		"iss": baseURL,
		// "jti": (String identifier for the token) not defined nor required
	}

	if introspectType == models.AccessToken {
		introspectionData["token_type"] = "Bearer"
	}

	notifier := ioc.GetNotifier(c)
	go notifier.Announce("client.token_introspection", utils.H{
		"Email":      shared.GetUserDefaultEmailForNotifications(c, user),
		"FirstName":  user.FirstName,
		"ClientName": client.Name,
		"CreatedAt":  time.Now().UTC().Format(time.RFC850),
	})

	c.JSON(http.StatusOK, introspectionData)
}
