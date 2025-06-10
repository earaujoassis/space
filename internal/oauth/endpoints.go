package oauth

import (
	"github.com/gin-gonic/gin"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the OAuth endpoints scope
func ExposeRoutes(router *gin.Engine) {
	endpoints := router.Group("/")
	{
		endpoints.GET("/.well-known/oauth-authorization-server", getOAuthAuthorizationServer)
		endpoints.GET("/oauth/authorize", authorizeHandler)
		endpoints.POST("/oauth/authorize", authorizeHandler)
		endpoints.POST("/oauth/token", tokenHandler)
		endpoints.POST("/oauth/revoke", revokeHandler)
		endpoints.POST("/oauth/introspect", introspectHandler)
		endpoints.GET("/authorize", authorizeHandler)
		endpoints.POST("/authorize", authorizeHandler)
		endpoints.POST("/token", tokenHandler)
		endpoints.POST("/revoke", revokeHandler)
		endpoints.POST("/introspect", introspectHandler)
	}
}
