package api

import (
	"github.com/gin-gonic/gin"
)

// exposeClientsRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the clients resource
func exposeClientsRoutes(router *gin.RouterGroup) {
	// In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
	// TODO Improve security for this endpoint avoiding any overhead
	router.GET("/clients/:client_id/credentials", requiresApplicationSession(), clientsCredentialsHandler)

	clientsRoutes := router.Group("/clients")
	clientsRoutes.Use(requiresConformance())
	clientsRoutes.Use(actionTokenBearerAuthorization())
	clientsRoutes.Use(requiresApplicationSession())
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		clientsRoutes.GET("", clientsListHandler)

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		clientsRoutes.POST("/create", clientsCreateHandler)

		// In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
		// Authorization type: action token / Bearer (for web use)
		// TODO Improve security for this endpoint avoiding any overhead
		clientsRoutes.PATCH("/:client_id/profile", clientsProfileHandler)
	}
}
