package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/utils"
)

func profileEmailVerificationHandler(c *gin.Context) {
	authorizationBearer := c.Query("_")
	repositories := ioc.GetRepositories(c)
	action := repositories.Actions().Authentication(authorizationBearer)
	user := repositories.Users().FindByID(action.UserID)
	if action.UUID == "" || !action.GrantsWriteAbility() || !action.CanUpdateUser() {
		c.HTML(http.StatusUnauthorized, "error.email_confirmation", utils.H{
			"Title":    " - Email Confirmation",
			"Internal": true,
		})
		return
	}

	user.EmailVerified = true

	repositories.Users().Save(&user)
	repositories.Actions().Delete(action)
	c.Redirect(http.StatusFound, "/")
}
