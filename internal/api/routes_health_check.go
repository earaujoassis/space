package api

import (
	"github.com/gin-gonic/gin"
)

func exposeHealthCheck(router *gin.RouterGroup) {
	healthCheckRoutes := router.Group("/")
	{
		healthCheckRoutes.GET("/health-check", healthCheckHandler)
	}
}
