package api

import (
	"github.com/gin-gonic/gin"
)

// exposeSessionsRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the sessions resource
func exposeSessionsRoutes(router *gin.RouterGroup) {
	sessionsRoutes := router.Group("/sessions")
	sessionsRoutes.Use(requiresConformance())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		sessionsRoutes.POST("/create", sessionsCreateHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		sessionsRoutes.POST("/magic", sessionsMagicHandler)
	}
}
