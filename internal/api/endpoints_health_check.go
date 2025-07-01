package api

import (
	"github.com/gin-gonic/gin"
)

func exposeHealthCheck(router *gin.RouterGroup) {
	healthCheckRoutes := router.Group("/health-check")
	{
		healthCheckRoutes.GET("", healthCheckHandler)
	}
}
