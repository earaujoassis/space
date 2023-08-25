package tasks

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"

	"github.com/earaujoassis/space/internal/api"
	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/datastore"
	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/utils"
	"github.com/earaujoassis/space/internal/web"
)

// Server is used to start and serve the application (REST API + Web front-end)
func Server() {
	datastore.Start()
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		defer func(c *gin.Context) {
			// TODO It is not displaying/logging the error
			if rec := recover(); rec != nil {
				if utils.MustServeJSON(c.Request.URL.Path, c.Request.Header.Get("Accept")) {
					c.JSON(http.StatusInternalServerError, utils.H{
						"_status":  "error",
						"_message": "Bad server",
						"error":    "The server found an error; aborting",
					})
				} else {
					c.HTML(http.StatusInternalServerError, "error.internal", utils.H{
						"Title":    " - Bad Server",
						"Internal": true,
					})
				}
			}
		}(c)
		c.Next()
	})
	router.NoRoute(func(c *gin.Context) {
		if utils.MustServeJSON(c.Request.URL.Path, c.Request.Header.Get("Accept")) {
			c.JSON(http.StatusNotFound, utils.H{
				"_status":  "error",
				"_message": "Not found",
				"error":    "Resource path not found",
			})
		} else {
			c.HTML(http.StatusNotFound, "error.not_found", utils.H{
				"Title":    " - Resource Not Found",
				"Internal": true,
			})
		}
	})
	web.ExposeRoutes(router)
	restAPI := router.Group("/api")
	api.ExposeRoutes(restAPI)
	router.Run(fmt.Sprintf(":%v", config.GetEnvVar("PORT")))
}

// ToggleFeature is used to enable or disable a feature-gate
func ToggleFeature() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Feature key: ")
	featureKey, _ := reader.ReadString('\n')
	featureKey = strings.Trim(featureKey, "\n")
	if feature.IsActive(featureKey) {
		feature.Disable(featureKey)
		fmt.Printf("Key `%s` is disabled\n", featureKey)
	} else {
		feature.Enable(featureKey)
		fmt.Printf("Key `%s` is enabled\n", featureKey)
	}
}
