package services

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api/helpers"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the services resource
func ExposeRoutes(router *gin.RouterGroup) {
	servicesRoutes := router.Group("/services")
	servicesRoutes.Use(helpers.RequiresConformance())
	servicesRoutes.Use(helpers.RequiresApplicationSession())
	servicesRoutes.Use(helpers.ActionTokenBearerAuthorization())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.GET("", listHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.POST("", createHandler)
	}
}
