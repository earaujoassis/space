package tasks

import (
    "fmt"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/web"
    "github.com/earaujoassis/space/api"
)

func Server() {
    datastore.Start()
    router := gin.Default()
    web.ExposeRoutes(router)
    restApi := router.Group("/api")
    api.ExposeRoutes(restApi)
    router.Run(fmt.Sprintf(":%v", config.GetConfig("PORT")))
}
