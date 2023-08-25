package api

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/internal/services"
)

func exposeHealthCheck(router *gin.RouterGroup) {
    healthCheckRoutes := router.Group("/")
    {
        healthCheckRoutes.GET("/health-check", func(c *gin.Context) {
            if services.IsDatastoreConnectedAndHealthy() {
                c.String(http.StatusOK, "healthy")
            } else {
                c.String(http.StatusOK, "unhealthy")
            }
        })
    }
}
