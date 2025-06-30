package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

const (
	passwordType      = "password"
	secretsType       = "secrets"
	emailVerification = "email_verification"
)

func usersUpdateRequestHandler(c *gin.Context) {
	var holder = c.PostForm("holder")
	var requestType = c.PostForm("request_type")
	var host = fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)

	if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "update request was not created",
			"error":    "must use valid holder string",
		})
		return
	}

	repositories := ioc.GetRepositories(c)
	switch requestType {
	case passwordType:
		user := repositories.Users().FindByAccountHolder(holder)
		client := repositories.Clients().FindOrCreate(models.DefaultClient)
		if user.IsSavedRecord() {
			actionToken := models.Action{
				User:        user,
				Client:      client,
				IP:          c.Request.RemoteAddr,
				UserAgent:   c.Request.UserAgent(),
				Scopes:      models.WriteScope,
				Description: models.UpdateUserAction,
			}
			repositories.Actions().Create(&actionToken)
			notifier := ioc.GetNotifier(c)
			go notifier.Announce("session.magic", utils.H{
				"Email":     user.Email,
				"FirstName": user.FirstName,
				"Callback":  fmt.Sprintf("%s/profile/password?_=%s", host, actionToken.Token),
			})
		}
	case secretsType:
		user := repositories.Users().FindByAccountHolder(holder)
		client := repositories.Clients().FindOrCreate(models.DefaultClient)
		if user.IsSavedRecord() {
			actionToken := models.Action{
				User:        user,
				Client:      client,
				IP:          c.Request.RemoteAddr,
				UserAgent:   c.Request.UserAgent(),
				Scopes:      models.WriteScope,
				Description: models.UpdateUserAction,
			}
			repositories.Actions().Create(&actionToken)
			notifier := ioc.GetNotifier(c)
			go notifier.Announce("session.magic", utils.H{
				"Email":     user.Email,
				"FirstName": user.FirstName,
				"Callback":  fmt.Sprintf("%s/profile/secrets?_=%s", host, actionToken.Token),
			})
		}
	case emailVerification:
		user := repositories.Users().FindByAccountHolder(holder)
		client := repositories.Clients().FindOrCreate(models.DefaultClient)
		if user.IsSavedRecord() {
			actionToken := models.Action{
				User:        user,
				Client:      client,
				IP:          c.Request.RemoteAddr,
				UserAgent:   c.Request.UserAgent(),
				Scopes:      models.WriteScope,
				Description: models.UpdateUserAction,
			}
			repositories.Actions().Create(&actionToken)
			notifier := ioc.GetNotifier(c)
			go notifier.Announce("user.update.email_verification", utils.H{
				"Email":     user.Email,
				"FirstName": user.FirstName,
				"Callback":  fmt.Sprintf("%s/profile/email_verification?_=%s", host, actionToken.Token),
			})
		}
	default:
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "update request was not created",
			"error":    "request type not available",
		})
		return
	}

	// No Content is the default response for valid requests
	c.JSON(http.StatusNoContent, nil)
}
