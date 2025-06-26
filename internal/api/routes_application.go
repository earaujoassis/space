package api

import (
	"github.com/gin-gonic/gin"
)

// exposeApplicationRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	it represents API calls related to the UI application
//
//	in the REST API scope, for the application resource
func exposeApplicationRoutes(router *gin.RouterGroup) {
	applicationRoutes := router.Group("/application")
	applicationRoutes.Use(requiresConformance())
	applicationRoutes.Use(requiresApplicationSession())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		applicationRoutes.GET("/bootstrap", applicationBootstrapHandler)
	}
}
