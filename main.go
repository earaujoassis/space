package main

import (
    "fmt"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/web"
    "github.com/earaujoassis/space/authentication"
)

func main() {
    router := gin.Default()
    web.ExposeRoutes(router)
    restApi := router.Group("/api/v1")
    authentication.ExposeRoutes(restApi)
    router.Run(fmt.Sprintf(":%v", config.GetConfig("http.port")))
}
