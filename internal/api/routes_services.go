package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// exposeServicesRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the services resource
func exposeServicesRoutes(router *gin.RouterGroup) {
	// Requires X-Requested-By and Origin (same-origin policy)
	// Authorization type: action token / Bearer (for web use)
	router.GET("/services", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
		repositories := ioc.GetRepositories(c)
		action := c.MustGet("Action").(models.Action)
		session := sessions.Default(c)
		userPublicID := session.Get("user_public_id")
		user := repositories.Users().FindByPublicID(userPublicID.(string))
		if userPublicID == nil || user.IsNewRecord() || user.ID != action.UserID || !user.Admin {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
			c.JSON(http.StatusUnauthorized, utils.H{
				"_status":  "error",
				"_message": "Services are not available",
				"error":    shared.AccessDenied,
			})
			return
		}

		c.JSON(http.StatusOK, utils.H{
			"_status":  "success",
			"_message": "Services are available",
			"services": repositories.Services().GetAll(),
		})
	})

	servicesRoutes := router.Group("/services")
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.POST("/create", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			repositories := ioc.GetRepositories(c)
			session := sessions.Default(c)
			action := c.MustGet("Action").(models.Action)
			userPublicID := session.Get("user_public_id")
			user := repositories.Users().FindByPublicID(userPublicID.(string))
			if userPublicID == nil || user.IsNewRecord() || user.ID != action.UserID || !user.Admin {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "Service was not created",
					"error":    shared.AccessDenied,
				})
				return
			}

			serviceName := c.PostForm("name")
			serviceDescription := c.PostForm("description")
			canonicalURI := c.PostForm("canonical_uri")
			logoURI := c.PostForm("logo_uri")

			service := repositories.Services().Create(serviceName,
				serviceDescription,
				canonicalURI,
				logoURI)

			if service.IsNewRecord() {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "Service was not created",
					"error":    "cannot create Service",
					"service":  service,
				})
			} else {
				c.JSON(http.StatusNoContent, nil)
			}
		})
	}
}
