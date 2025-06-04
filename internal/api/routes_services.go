package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/oauth"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

// exposeServicesRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API escope, for the services resource
func exposeServicesRoutes(router *gin.RouterGroup) {
	servicesRoutes := router.Group("/services")
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.GET("/", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			action := c.MustGet("Action").(models.Action)
			session := sessions.Default(c)
			userPublicID := session.Get("userPublicID")
			user := services.FindUserByPublicID(userPublicID.(string))
			if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || !user.Admin {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "Services are not available",
					"error":    oauth.AccessDenied,
				})
				return
			}

			c.JSON(http.StatusOK, utils.H{
				"_status":  "success",
				"_message": "Services are available",
				"services":  services.Services(),
			})
		})

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.POST("/create", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			session := sessions.Default(c)
			action := c.MustGet("Action").(models.Action)
			userPublicID := session.Get("userPublicID")
			user := services.FindUserByPublicID(userPublicID.(string))
			if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || !user.Admin {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "Service was not created",
					"error":    oauth.AccessDenied,
				})
				return
			}

			serviceName := c.PostForm("name")
			serviceDescription := c.PostForm("description")
			canonicalURI := c.PostForm("canonical_uri")
			logoURI := c.PostForm("logo_uri")

			service := services.CreateNewService(serviceName,
				serviceDescription,
				canonicalURI,
				logoURI)

			if service.ID == 0 {
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
