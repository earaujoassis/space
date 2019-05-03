package tasks

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/web"
    "github.com/earaujoassis/space/api"
    "github.com/earaujoassis/space/feature"
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

// ToggleFeature is used to enable or disable a feature-gate
func ToggleFeature() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Feature key: ")
    featureKey, _ := reader.ReadString('\n')
    featureKey = strings.Trim(featureKey, "\n")
    if feature.IsActive(featureKey) {
        feature.Disable(featureKey)
        fmt.Printf("Key `%s` was disabled\n", featureKey)
    } else {
        feature.Enable(featureKey)
        fmt.Printf("Key `%s` was enabled\n", featureKey)
    }
}
