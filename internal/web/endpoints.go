package web

import (
	"github.com/gin-gonic/gin"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the Web scope
func ExposeRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("web/templates/*.html")
	router.HTMLRender = createCustomRender()
	router.Static("/public", "web/public")
	views := router.Group("/")
	{
		views.GET("/", satelliteHandler)
		views.GET("/applications", satelliteHandler)
		views.GET("/clients", satelliteHandler)
		views.GET("/clients/edit", satelliteHandler)
		views.GET("/clients/edit/scopes", satelliteHandler)
		views.GET("/clients/new", satelliteHandler)
		views.GET("/services", satelliteHandler)
		views.GET("/services/new", satelliteHandler)
		views.GET("/notifications", satelliteHandler)
		views.GET("/profile", satelliteHandler)
		views.GET("/security", satelliteHandler)
		views.GET("/profile/password", profilePasswordHandler)
		views.GET("/profile/secrets", profileSecretsHandler)
		views.GET("/profile/email_verification", profileEmailVerificationHandler)
		views.GET("/signup", signupHandler)
		views.GET("/signin", signinHandler)
		views.GET("/signout", signoutHandler)
		views.GET("/session", sessionHandler)
		views.GET("/error", errorHandler)
	}
}
