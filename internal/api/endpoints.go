package api

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api/clients"
	"github.com/earaujoassis/space/internal/api/health_check"
	"github.com/earaujoassis/space/internal/api/self"
	"github.com/earaujoassis/space/internal/api/services"
	"github.com/earaujoassis/space/internal/api/sessions"
	"github.com/earaujoassis/space/internal/api/users"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope
func ExposeRoutes(router *gin.Engine) {
	restAPI := router.Group("/api")

	health_check.ExposeRoutes(restAPI)
	users.ExposeRoutes(restAPI)
	self.ExposeRoutes(restAPI)
	sessions.ExposeRoutes(restAPI)
	clients.ExposeRoutes(restAPI)
	services.ExposeRoutes(restAPI)
}
