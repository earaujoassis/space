package tasks

import (
    "fmt"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/web"
    "github.com/earaujoassis/space/api"
)

// Server is used to start and serve the application (REST API + Web front-end)
func Server() {
    datastore.Start()
    router := gin.Default()
    web.ExposeRoutes(router)
    restAPI := router.Group("/api")
    api.ExposeRoutes(restAPI)
    router.Run(fmt.Sprintf(":%v", config.GetConfig("PORT")))
}
