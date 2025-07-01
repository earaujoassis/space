package admin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func requiresAdminApplicationSession() gin.HandlerFunc {
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
		if applicationSession.IsSavedRecord() && applicationSession.User.Admin {
			c.Set("User", applicationSession.User)
			c.Next()
			return
		} else if applicationSession.IsSavedRecord() {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}
		c.Redirect(http.StatusFound, "/signin")
		c.Abort()
	}
}
