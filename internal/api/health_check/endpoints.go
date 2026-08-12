package health_check

import (
	"github.com/gin-gonic/gin"
)

func ExposeRoutes(router *gin.RouterGroup) {
	router.GET("/health-check", handler)
}
