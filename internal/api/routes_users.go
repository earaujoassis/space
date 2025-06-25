package api

import (
	"github.com/gin-gonic/gin"
)

// exposeUsersRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the users resource
func exposeUsersRoutes(router *gin.RouterGroup) {
	usersRoutes := router.Group("/users")
	usersRoutes.Use(requiresConformance())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		usersRoutes.POST("/create", usersCreateHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersRoutes.POST("/update/request", usersUpdateRequestHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersRoutes.PATCH("/update/password", usersUpdatePasswordHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.PATCH("/update/adminify", actionTokenBearerAuthorization(), usersUpdateAdminifyHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/profile", actionTokenBearerAuthorization(), usersProfileHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/clients", actionTokenBearerAuthorization(), usersClientsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.DELETE("/:user_id/clients/:client_id/revoke", actionTokenBearerAuthorization(), usersClientsRevokeHandler)
	}
}
