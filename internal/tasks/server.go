package tasks

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/admin"
	"github.com/earaujoassis/space/internal/api"
	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/oauth"
	"github.com/earaujoassis/space/internal/oidc"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
	"github.com/earaujoassis/space/internal/web"
)

// Server is used to start and serve the application components:
//
//	REST API
//	Web Components & Views
//	Admin & Monitoring Views
func Server(cfg *config.Config) {
	router := SetupRouter(cfg)
	router.Run(fmt.Sprintf(":%v", cfg.GetEnvVar("PORT")))
}

func SetupRouter(cfg *config.Config) *gin.Engine {
	appCtx, err := ioc.NewAppContext(cfg)
	if err != nil {
		logs.Propagatef(logs.Panic, "Could not create app context: %s\n", err)
	}
	gin.DisableConsoleColor()
	router := gin.Default()
	security.SetTrustedProxies(router)
	router.Use(ioc.InjectAppContext(appCtx))
	router.RedirectTrailingSlash = false
	store := cookie.NewStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{
		Secure:   (cfg.IsEnvironment("production") && cfg.SessionSecure),
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("space.session", store))
	router.Use(func(c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {
				defer logs.Propagatef(logs.Error, "%+v\n%s\n", fmt.Errorf("%v", rec), string(debug.Stack()))
				if shared.MustServeJSON(c.Request.URL.Path, c.Request.Header.Get("Accept")) {
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
		if shared.MustServeJSON(c.Request.URL.Path, c.Request.Header.Get("Accept")) {
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
	oauth.ExposeRoutes(router)
	oidc.ExposeRoutes(router)
	api.ExposeRoutes(router)
	web.ExposeRoutes(router)
	admin.ExposeRoutes(router, cfg)

	return router
}
