package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/policy"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/services/communications"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// exposeSessionsRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the sessions resource
func exposeSessionsRoutes(router *gin.RouterGroup) {
	sessionsRoutes := router.Group("/sessions")
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		sessionsRoutes.POST("/create", requiresConformance, func(c *gin.Context) {
			var holder = c.PostForm("holder")
			var state = c.PostForm("state")

			var IP = c.Request.RemoteAddr
			var userID = IP
			var statusSignInAttempts = policy.SignInAttemptStatus(IP)

			if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "Session was not created",
					"error":    "must use valid holder string",
				})
				return
			}

			user := services.FindUserByAccountHolder(holder)
			client := services.FindOrCreateClient(services.DefaultClient)
			if user.ID != 0 && statusSignInAttempts != policy.Blocked {
				userID = user.UUID
				statusSignInAttempts = policy.SignInAttemptStatus(userID)
				if user.Authentic(c.PostForm("password"), c.PostForm("passcode")) && statusSignInAttempts != policy.Blocked {
					session := services.CreateSession(user, client,
						c.Request.RemoteAddr,
						c.Request.UserAgent(),
						models.PublicScope,
						models.GrantToken)
					if session.ID != 0 {
						go communications.Announce("session.created", utils.H{
							"Email":     user.Email,
							"FirstName": user.FirstName,
							"IP":        session.IP,
							"CreatedAt": session.CreatedAt.Format(time.RFC850),
						})
						policy.RegisterSuccessfulSignIn(user.UUID)
						policy.RegisterSuccessfulSignIn(IP)
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
			policy.RegisterSignInAttempt(userID)
			c.JSON(http.StatusBadRequest, utils.H{
				"_status":  "error",
				"_message": "Session was not created",
				"error":    shared.AccessDenied,
				"attempts": statusSignInAttempts,
			})
		})

		// Requires X-Requested-By and Origin (same-origin policy)
		sessionsRoutes.POST("/magic", requiresConformance, func(c *gin.Context) {
			var holder = c.PostForm("holder")
			var next = c.PostForm("next")
			var state = c.PostForm("state")

			var host = fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)

			var IP = c.Request.RemoteAddr
			var userID = IP
			var statusSignInAttempts = policy.SignInAttemptStatus(IP)

			if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "Magic Session was not created",
					"error":    "must use valid holder string",
				})
				return
			}

			user := services.FindUserByAccountHolder(holder)
			client := services.FindOrCreateClient(services.DefaultClient)
			if user.ID != 0 && statusSignInAttempts != policy.Blocked {
				userID = user.UUID
				statusSignInAttempts = policy.SignInAttemptStatus(userID)
				if statusSignInAttempts != policy.Blocked {
					session := services.CreateSession(user, client,
						c.Request.RemoteAddr,
						c.Request.UserAgent(),
						models.PublicScope,
						models.GrantToken)
					if session.ID != 0 {
						go communications.Announce("session.magic", utils.H{
							"Email":     user.Email,
							"FirstName": user.FirstName,
							"CreatedAt": session.CreatedAt.Format(time.RFC850),
							"Callback": fmt.Sprintf("%s/session?client_id=%s&code=%s&grant_type=authorization_code&scope=%s&state=%s&_=%s",
								host, client.Key, session.Token, session.Scopes, state, next),
						})
						policy.RegisterSuccessfulSignIn(user.UUID)
						policy.RegisterSuccessfulSignIn(IP)
						c.JSON(http.StatusNoContent, nil)
						return
					}
				}
			}
			policy.RegisterSignInAttempt(userID)
			c.JSON(http.StatusNoContent, nil)
		})
	}
}
