package tasks

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/api"
	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/datastore"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/utils"
	"github.com/earaujoassis/space/internal/web"
	"github.com/earaujoassis/space/internal/oauth"
)

// Server is used to start and serve the application (REST API + Web front-end)
func Server() {
	datastore.InitConnection()
	gin.DisableConsoleColor()
	router := gin.Default()
	cfg := config.GetGlobalConfig()
	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{
		Secure:   (config.IsEnvironment("production") && cfg.SessionSecure),
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("space.session", store))
	router.Use(func(c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {
				defer logs.Propagatef(logs.Error, "%+v\n%s\n", errors.New(fmt.Sprintf("%v", rec)), string(debug.Stack()))
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
	oauth.ExposeRoutes(router)
	restAPI := router.Group("/api")
	api.ExposeRoutes(restAPI)
	router.Run(fmt.Sprintf(":%v", config.GetEnvVar("PORT")))
}
