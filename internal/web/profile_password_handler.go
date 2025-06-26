package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/utils"
)

func profilePasswordHandler(c *gin.Context) {
	authorizationBearer := c.Query("_")
	repositories := ioc.GetRepositories(c)
	action := repositories.Actions().Authentication(authorizationBearer)
	if action.UUID == "" || !action.GrantsWriteAbility() || !action.CanUpdateUser() {
		c.HTML(http.StatusUnauthorized, "error.password_update", utils.H{
			"Title":    " - Update Resource Owner Credential",
			"Internal": true,
		})
		return
	}

	c.HTML(http.StatusOK, "satellite", utils.H{
		"Title":     " - Update Resource Owner Credential",
		"Satellite": "amalthea",
		"Internal":  true,
		"Data": utils.H{
			"action_token": action.Token,
		},
	})
}
