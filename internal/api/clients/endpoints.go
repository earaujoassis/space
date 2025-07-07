package clients

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api/helpers"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the clients resource
func ExposeRoutes(router *gin.RouterGroup) {
	// In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
	// TODO Improve security for this endpoint avoiding any overhead
	router.GET("/clients/:client_id/credentials",
		helpers.RequiresApplicationSession(),
		credentialsHandler)

	clientsRoutes := router.Group("/clients")
	clientsRoutes.Use(helpers.RequiresConformance())
	clientsRoutes.Use(helpers.RequiresApplicationSession())
	clientsRoutes.Use(helpers.ActionTokenBearerAuthorization())
	clientsRoutes.Use(helpers.RequireActionTokenFromAuthenticatedUser())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		clientsRoutes.GET("", listHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		clientsRoutes.POST("", createHandler)

		// In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
		// Authorization type: action token / Bearer (for web use)
		// TODO Improve security for this endpoint avoiding any overhead
		clientsRoutes.PATCH("/:client_id/profile", profileHandler)
	}
}
