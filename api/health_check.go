package api

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/datastore"
)

func exposeHealthCheck(router *gin.RouterGroup) {
    healthCheckRoutes := router.Group("/")
    {
        healthCheckRoutes.GET("/health-check", func(c *gin.Context) {
            var count struct{
                Count int64
            }

            dataStoreSession := datastore.GetDataStoreConnection()
            dataStoreSession.
                Raw("SELECT count(*) AS count FROM clients;").
                Scan(&count)
            if count.Count >= 0 {
                c.String(http.StatusOK, "healthy")
            } else {
                c.String(http.StatusOK, "unhealthy")
            }
        })
    }
}
