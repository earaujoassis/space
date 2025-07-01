package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"

	"github.com/earaujoassis/space/internal/config"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the Admin scope
func ExposeRoutes(router *gin.Engine, cfg *config.Config) {
	externalRoutes := router.Group("/_")
	externalRoutes.Use(requiresAdminApplicationSession())

	redisAddr := fmt.Sprintf("%s:%d", cfg.MemorystoreHost, cfg.MemorystorePort)
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/_/monitoring",
		RedisConnOpt: asynq.RedisClientOpt{Addr: redisAddr, DB: cfg.MemorystoreIndex},
	})

	externalRoutes.Any("/monitoring", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	externalRoutes.Any("/monitoring/*any", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}
