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
		usersRoutes.POST("", usersCreateHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/profile",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersProfileHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/clients",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersClientsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.DELETE("/:user_id/clients/:client_id/revoke",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersClientsRevokeHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/sessions",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersSessionsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.DELETE("/:user_id/sessions/:session_id/revoke",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersSessionsRevokeHandler)
	}
	usersMeRoutes := usersRoutes.Group("/me")
	usersMeRoutes.Use(requiresConformance())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.POST("/requests", usersMeRequestsHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.GET("/workspace",
			requiresApplicationSession(),
			usersMeWorkspaceHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.PATCH("/password", usersMePasswordHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersMeRoutes.PATCH("/admin",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersMeAdminHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.GET("/emails",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersMeEmailsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.POST("/emails",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersMeEmailsCreateHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.GET("/settings",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersMeSettingsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		usersMeRoutes.PATCH("/settings",
			requiresApplicationSession(),
			actionTokenBearerAuthorization(),
			usersMeSettingsPatchHandler)
	}
}
