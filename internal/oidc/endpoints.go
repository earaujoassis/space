package oidc

import (
	"github.com/gin-gonic/gin"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the OpenID Connect endpoints scope
func ExposeRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("internal/oidc/templates/*.html")
	router.HTMLRender = createCustomRender()
	endpoints := router.Group("/")
	{
		endpoints.GET("/.well-known/openid-configuration", getOpenIdConfiguration)
		endpoints.GET("/oidc/userinfo", userinfoHandler)
		endpoints.GET("/oidc/jwks", jwksHandler)
		endpoints.GET("/oidc/authorize", authorizeHandler)
		endpoints.POST("/oidc/authorize", authorizeHandler)
	}
}
