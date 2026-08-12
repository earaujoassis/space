package sessions

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api/helpers"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the sessions resource
func ExposeRoutes(router *gin.RouterGroup) {
	sessionsRoutes := router.Group("/sessions")
	sessionsRoutes.Use(helpers.RequiresConformance())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		sessionsRoutes.POST("", createHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		sessionsRoutes.POST("/requests", requestsHandler)
	}
}
