package api

import (
	"github.com/gin-gonic/gin"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope
func ExposeRoutes(router *gin.Engine) {
	restAPI := router.Group("/api")

	exposeUsersRoutes(restAPI)
	exposeSessionsRoutes(restAPI)
	exposeClientsRoutes(restAPI)
	exposeServicesRoutes(restAPI)
	exposeHealthCheck(restAPI)
}
