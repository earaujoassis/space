package api

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

const (
	passwordType          = "password"
	secretsType           = "secrets"
	emailVerificationType = "email_verification"
)

func usersMeRequestsHandler(c *gin.Context) {
	requestType := c.PostForm("request_type")
	availableTypes := []string{passwordType, secretsType, emailVerificationType}
	if !slices.Contains(availableTypes, requestType) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "update request was not created",
			"error":    "request type not available",
		})
		return
	}

	holder := c.PostForm("holder")
	if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "update request was not created",
			"error":    "must use valid holder field",
		})
		return
	}

	repositories := ioc.GetRepositories(c)
	client := repositories.Clients().FindOrCreate(models.DefaultClient)
	user := repositories.Users().FindByAccountHolder(holder)
	if user.IsNewRecord() {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "update request was not created",
			"error":    "must use valid holder field",
		})
		return
	}

	host := fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)
	switch requestType {
	case passwordType:
		actionToken := models.Action{
			User:        user,
			Client:      client,
			IP:          c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
			Scopes:      models.WriteScope,
			Description: models.UpdateUserAction,
		}
		repositories.Actions().Create(&actionToken)
		notifier := ioc.GetNotifier(c)
		go notifier.Announce("user.update_password", utils.H{
			"Email":     user.Email,
			"FirstName": user.FirstName,
			"Callback":  fmt.Sprintf("%s/profile/password?_=%s", host, actionToken.Token),
		})
	case secretsType:
		actionToken := models.Action{
			User:        user,
			Client:      client,
			IP:          c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
			Scopes:      models.WriteScope,
			Description: models.UpdateUserAction,
		}
		repositories.Actions().Create(&actionToken)
		notifier := ioc.GetNotifier(c)
		go notifier.Announce("user.update_secrets", utils.H{
			"Email":     user.Email,
			"FirstName": user.FirstName,
			"Callback":  fmt.Sprintf("%s/profile/secrets?_=%s", host, actionToken.Token),
		})
	case emailVerificationType:
		var emailAddress = c.PostForm("email")
		if !security.ValidEmail(emailAddress) || !repositories.Users().HoldsEmail(user, emailAddress) {
			c.JSON(http.StatusBadRequest, utils.H{
				"_status":  "error",
				"_message": "update request was not created",
				"error":    "must use valid email field",
			})
			return
		}
		actionToken := models.Action{
			User:        user,
			Client:      client,
			IP:          c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
			Scopes:      models.WriteScope,
			Description: models.UpdateUserAction,
			Payload:     emailAddress,
		}
		repositories.Actions().Create(&actionToken)
		notifier := ioc.GetNotifier(c)
		go notifier.Announce("user.email_verification", utils.H{
			"Email":     emailAddress,
			"FirstName": user.FirstName,
			"Callback":  fmt.Sprintf("%s/profile/email_verification?_=%s", host, actionToken.Token),
		})
	}

	// No Content is the default response for valid requests
	c.JSON(http.StatusNoContent, nil)
}
