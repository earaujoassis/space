package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func applicationBootstrapHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	fg := ioc.GetFeatureGate(c)
	user := c.MustGet("User").(models.User)
	client := repositories.Clients().FindOrCreate(models.DefaultClient)

	actionToken := models.Action{
		User:        user,
		Client:      client,
		IP:          c.ClientIP(),
		UserAgent:   c.Request.UserAgent(),
		Scopes:      models.WriteScope,
		Description: models.NotSpecialAction,
	}
	err := repositories.Actions().Create(&actionToken)
	if err == nil {
		c.JSON(http.StatusOK, utils.H{
			"application": utils.H{
				"action_token":  actionToken.Token,
				"user_id":       user.UUID,
				"user_is_admin": user.Admin,
				"feature.gates": utils.H{
					"user.adminify": fg.IsActive("user.adminify"),
				},
			},
		})
	} else {
		c.JSON(http.StatusInternalServerError, utils.H{
			"error": shared.InternalError,
		})
	}
}
