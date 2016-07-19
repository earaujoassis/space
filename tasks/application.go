package tasks

import (
    "fmt"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/web"
    "github.com/earaujoassis/space/authentication"
)

func Server(publicFolder, templateFolder string) {
    datastore.Start()
    config.SetConfig("template_folder", templateFolder)
    config.SetConfig("public_folder", publicFolder)
    router := gin.Default()
    web.ExposeRoutes(router)
    restApi := router.Group("/api/v1")
    authentication.ExposeRoutes(restApi)
    router.Run(fmt.Sprintf(":%v", config.GetConfig("http.port")))
}
