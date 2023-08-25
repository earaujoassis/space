package api

import (
	"github.com/gin-gonic/gin"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API escope
func ExposeRoutes(router *gin.RouterGroup) {
	exposeUsersRoutes(router)
	exposeSessionsRoutes(router)
	exposeClientsRoutes(router)
	exposeHealthCheck(router)
}
