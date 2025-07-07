package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func RequiresConformance() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)
		correctXRequestedBy := c.Request.Header.Get("X-Requested-By") == "SpaceApi"
		// WARNING The Origin header attribute sometimes is not sent; we should not block these requests
		sameOriginPolicy := c.Request.Header.Get("Origin") == "" || host == c.Request.Header.Get("Origin")
		if correctXRequestedBy && sameOriginPolicy {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, utils.H{
				"error": "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy",
			})
			c.Abort()
		}
	}
}

func RequiresApplicationSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		var applicationSession models.Session

		session := sessions.Default(c)
		repositories := ioc.GetRepositories(c)
		applicationTokenInterface := session.Get(shared.CookieSessionKey)
		applicationToken := utils.StringValue(applicationTokenInterface)
		if applicationToken != "" && !security.ValidToken(applicationToken) {
			applicationSession = models.Session{}
		} else {
			applicationSession = repositories.Sessions().FindByToken(applicationToken, models.ApplicationToken)
		}
		if applicationSession.IsSavedRecord() {
			c.Set("CurrentSession", applicationSession)
			c.Set("User", applicationSession.User)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "User must be authenticated",
			"error":    "unauthorized request",
		})
		c.Abort()
	}
}

// The following Authorization method is used by the web client, with an action token
func ActionTokenBearerAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationBearer := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		if !security.ValidToken(authorizationBearer) {
			c.JSON(http.StatusBadRequest, utils.H{
				"error": "must use valid token field",
			})
			c.Abort()
			return
		}

		repositories := ioc.GetRepositories(c)
		action := repositories.Actions().Authentication(authorizationBearer)
		if action.UUID == "" || !action.GrantsReadAbility() {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
			c.JSON(http.StatusUnauthorized, utils.H{
				"error": shared.AccessDenied,
			})
			c.Abort()
			return
		}

		c.Set("Action", action)
		c.Next()
	}
}

func RequireMatchBetweenActionTokenAndAuthenticatedUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		action := c.MustGet("Action").(models.Action)
		authenticatedUser := c.MustGet("User").(models.User)
		if authenticatedUser.ID != action.UserID {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
			c.JSON(http.StatusUnauthorized, utils.H{
				"_status":  "error",
				"_message": "Groups not available",
				"error":    shared.AccessDenied,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
