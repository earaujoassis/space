package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

func applicationBootstrapHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	fg := ioc.GetFeatureGate(c)
	session := sessions.Default(c)
	userPublicID := session.Get("user_public_id")
	if userPublicID == nil {
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Application bootstrap failed",
			"error":    "unauthorized application bootstrap",
		})
	}
	client := repositories.Clients().FindOrCreate(models.DefaultClient)
	user := repositories.Users().FindByPublicID(userPublicID.(string))

	actionToken := models.Action{
		User:        user,
		Client:      client,
		IP:          c.Request.RemoteAddr,
		UserAgent:   c.Request.UserAgent(),
		Scopes:      models.WriteScope,
		Description: models.NotSpecialAction,
	}
	repositories.Actions().Create(&actionToken)
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
}
