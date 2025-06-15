package tasks

import (
	"fmt"

	"github.com/earaujoassis/space/internal/config"
)

// Server is used to start and serve the application (REST API + Web front-end)
func Server(cfg *config.Config) {
	router := SetupRouter(cfg)
	router.Run(fmt.Sprintf(":%v", cfg.GetEnvVar("PORT")))
}
