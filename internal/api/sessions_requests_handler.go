package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/policy"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

const (
	passwordlessSigninType = "passwordless_signin"
)

func sessionsRequestsHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	rls := ioc.GetRateLimitService(c)

	var requestType = c.PostForm("request_type")
	switch requestType {
	case passwordlessSigninType:
		var holder = c.PostForm("holder")
		var next = c.PostForm("next")
		var state = c.PostForm("state")

		var host = fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)

		var IP = c.ClientIP()
		var userID = IP
		var statusSignInAttempts = rls.SignInAttemptStatus(IP)

		if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
			c.JSON(http.StatusBadRequest, utils.H{
				"_status":  "error",
				"_message": "Magic Session was not created",
				"error":    "must use valid holder field",
			})
			return
		}

		user := repositories.Users().FindByAccountHolder(holder)
		client := repositories.Clients().FindOrCreate(models.DefaultClient)
		if user.IsSavedRecord() && statusSignInAttempts != policy.Blocked {
			userID = user.UUID
			statusSignInAttempts = rls.SignInAttemptStatus(userID)
			if statusSignInAttempts != policy.Blocked {
				session := models.Session{
					User:      user,
					Client:    client,
					IP:        c.ClientIP(),
					UserAgent: c.Request.UserAgent(),
					Scopes:    models.PublicScope,
					TokenType: models.GrantToken,
				}
				repositories.Sessions().Create(&session)
				if session.IsSavedRecord() {
					notifier := ioc.GetNotifier(c)
					go notifier.Announce(user, "session.magic", utils.H{
						"Email":     user.Email,
						"FirstName": user.FirstName,
						"CreatedAt": time.Now().UTC().Format(time.RFC850),
						"Callback": fmt.Sprintf("%s/session?client_id=%s&code=%s&grant_type=authorization_code&scope=%s&state=%s&_=%s",
							host, client.Key, session.Token, session.Scopes, state, next),
					})
					rls.RegisterSuccessfulSignIn(user.UUID)
					rls.RegisterSuccessfulSignIn(IP)
					c.JSON(http.StatusNoContent, nil)
					return
				}
			}
		}
		rls.RegisterSignInAttempt(userID)
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
