package self

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api/helpers"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the users resource
func ExposeRoutes(router *gin.RouterGroup) {
	group := router.Group("/users/me")
	group.Use(helpers.RequiresConformance())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		group.POST("/requests", requestsHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		group.GET("/workspace",
			helpers.RequiresApplicationSession(),
			workspaceHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		group.PATCH("/password", passwordHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		group.PATCH("/admin",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			adminHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		group.GET("/emails",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			helpers.RequireActionTokenFromAuthenticatedUser(),
			emailsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		group.POST("/emails",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			helpers.RequireActionTokenFromAuthenticatedUser(),
			emailsCreateHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		group.GET("/settings",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			helpers.RequireActionTokenFromAuthenticatedUser(),
			settingsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		group.PATCH("/settings",
			helpers.RequiresApplicationSession(),
			helpers.ActionTokenBearerAuthorization(),
			helpers.RequireActionTokenFromAuthenticatedUser(),
			settingsPatchHandler)
	}
}
