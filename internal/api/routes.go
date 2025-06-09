package api

import (
	"github.com/gin-gonic/gin"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope
func ExposeRoutes(router *gin.RouterGroup) {
	exposeUsersRoutes(router)
	exposeSessionsRoutes(router)
	exposeClientsRoutes(router)
	exposeServicesRoutes(router)
	exposeHealthCheck(router)
	exposeApplicationRoutes(router)
}
