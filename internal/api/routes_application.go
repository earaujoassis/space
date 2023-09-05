package api

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

// exposeApplicationRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	it represents API calls related to the UI application
//
//	in the REST API escope, for the application resource
func exposeApplicationRoutes(router *gin.RouterGroup) {
	applicationRoutes := router.Group("/application")
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		applicationRoutes.GET("/bootstrap", requiresConformance, func(c *gin.Context) {
			session := sessions.Default(c)
			userPublicID := session.Get("userPublicID")
			if userPublicID == nil {
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "Application bootstrap failed",
					"error":    "unauthorized application bootstrap",
				})
			}
			client := services.FindOrCreateClient(services.DefaultClient)
			user := services.FindUserByPublicID(userPublicID.(string))
			actionToken := services.CreateAction(user, client,
				c.Request.RemoteAddr,
				c.Request.UserAgent(),
				models.ReadWriteScope,
				models.NotSpecialAction,
			)
			c.JSON(http.StatusOK, utils.H{
				"application": utils.H{
					"action_token":  actionToken.Token,
					"user_id":       user.UUID,
					"user_is_admin": user.Admin,
					"feature.gates": utils.H{
						"user.adminify": feature.IsActive("user.adminify"),
					},
				},
			})
		})
	}
}
