package api

import (
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

func sessionsCreateHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	rls := ioc.GetRateLimitService(c)

	var holder = c.PostForm("holder")
	var state = c.PostForm("state")

	var IP = c.ClientIP()
	var userID = IP
	var statusSignInAttempts = rls.SignInAttemptStatus(IP)

	if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Session was not created",
			"error":    "must use valid holder field",
		})
		return
	}

	user := repositories.Users().FindByAccountHolder(holder)
	client := repositories.Clients().FindOrCreate(models.DefaultClient)
	if user.IsSavedRecord() && statusSignInAttempts != policy.Blocked {
		userID = user.UUID
		statusSignInAttempts = rls.SignInAttemptStatus(userID)
		if repositories.Users().Authentic(user, c.PostForm("password"), c.PostForm("passcode")) && statusSignInAttempts != policy.Blocked {
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
				go notifier.Announce("session.created", utils.H{
					"Email":     shared.GetUserDefaultEmailForNotifications(c),
					"FirstName": user.FirstName,
					"CreatedAt": time.Now().UTC().Format(time.RFC850),
				})
				rls.RegisterSuccessfulSignIn(user.UUID)
				rls.RegisterSuccessfulSignIn(IP)
				c.JSON(http.StatusOK, utils.H{
					"_status":      "created",
					"_message":     "Session was created",
					"scope":        session.Scopes,
					"grant_type":   "authorization_code",
					"code":         session.Token,
					"redirect_uri": "/session",
					"client_id":    client.Key,
					"state":        state,
				})
				return
			}
		}
	}
	rls.RegisterSignInAttempt(userID)
	c.JSON(http.StatusBadRequest, utils.H{
		"_status":  "error",
		"_message": "Session was not created",
		"error":    shared.AccessDenied,
		"attempts": statusSignInAttempts,
	})
}
