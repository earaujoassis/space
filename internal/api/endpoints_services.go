package api

import (
	"github.com/gin-gonic/gin"
)

// exposeServicesRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the services resource
func exposeServicesRoutes(router *gin.RouterGroup) {
	servicesRoutes := router.Group("/services")
	servicesRoutes.Use(requiresConformance())
	servicesRoutes.Use(requiresApplicationSession())
	servicesRoutes.Use(actionTokenBearerAuthorization())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.GET("", servicesListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		servicesRoutes.POST("", servicesCreateHandler)
	}
}
