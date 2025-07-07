package users

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api/helpers"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the users resource
func ExposeRoutes(router *gin.RouterGroup) {
	usersRoutes := router.Group("/users")
	usersRoutes.Use(helpers.RequiresConformance())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		usersRoutes.POST("", createHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/profile",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			profileHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/clients",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			clientsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.DELETE("/:user_id/clients/:client_id/revoke",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			clientsRevokeHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/sessions",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			sessionsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.DELETE("/:user_id/sessions/:session_id/revoke",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			sessionsRevokeHandler)
	}
}
