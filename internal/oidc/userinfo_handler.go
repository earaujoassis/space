package oidc

import (
	"fmt"
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

func userinfoHandler(c *gin.Context) {
	var userinfo utils.H
	var user models.User
	var client models.Client
	var scope string

	authorizationHeader := c.GetHeader("Authorization")
	token := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	if token == "" {
		c.Header("WWW-Authenticate", "Bearer")
		c.JSON(http.StatusUnauthorized, "")
		return
	}

	repositories := ioc.GetRepositories(c)
	tokenType := identifyTokenType(token)
	switch tokenType {
	case shared.TokenTypeIDToken:
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer error=\"%s\"", shared.InvalidToken))
		c.JSON(http.StatusUnauthorized, utils.H{
			"error":             shared.InvalidToken,
			"error_description": "ID token not accepted, use access token",
		})
		return
	case shared.TokenTypeAccessToken:
		if !security.ValidToken(token) {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer error=\"%s\"", shared.InvalidToken))
			c.JSON(http.StatusUnauthorized, utils.H{
				"error":             shared.InvalidToken,
				"error_description": "The access token expired",
			})
			return
		}
		session := repositories.Sessions().FindByToken(token, models.AccessToken)
		if session.IsNewRecord() {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer error=\"%s\"", shared.InvalidToken))
			c.JSON(http.StatusUnauthorized, utils.H{
				"error":             shared.InvalidToken,
				"error_description": "The access token expired",
			})
			return
		}
		if !strings.Contains(session.Scopes, models.OpenIDScope) {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer error=\"%s\"", shared.InsufficientScope))
			c.JSON(http.StatusForbidden, utils.H{
				"error":             shared.InsufficientScope,
				"error_description": "The access token does not have the required scope",
			})
			return
		}
		user = session.User
		client = session.Client
		scope = session.Scopes
	}

	if !strings.Contains(scope, models.ProfileScope) {
		userinfo = utils.H{
			"sub": user.PublicID,
		}
	} else {
		userinfo = utils.H{
			"sub":                user.PublicID,
			"name":               fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			"given_name":         user.FirstName,
			"family_name":        user.LastName,
			"preferred_username": user.Username,
			"zoneinfo":           "UTC",
			"locale":             "en-US",
			"updated_at":         user.UpdatedAt.Unix(),
		}
	}
	if strings.Contains(scope, models.EmailScope) {
		userinfo["email"] = user.Email
		userinfo["email_verified"] = user.EmailVerified
	}

	notifier := ioc.GetNotifier(c)
	go notifier.Announce(user, "client.userinfo_introspection", shared.NotificationTemplateData(c, utils.H{
		"Email":      shared.GetUserDefaultEmailForNotifications(c, user),
		"FirstName":  user.FirstName,
		"ClientName": client.Name,
		"CreatedAt":  time.Now().UTC().Format(time.RFC850),
	}))

	c.JSON(http.StatusOK, userinfo)
}
