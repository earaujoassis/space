package redis

import (
	"fmt"

	"github.com/earaujoassis/space/internal/config"
)

func getRedisURL(cfg *config.Config) string {
	if cfg.IsEnvironment("production") {
		return fmt.Sprintf("redis://:%s@%s:%d/%d",
			cfg.MemorystorePassword,
			cfg.MemorystoreHost,
			cfg.MemorystorePort,
			cfg.MemorystoreIndex)
	} else {
		return fmt.Sprintf("redis://%s:%d/%d",
			cfg.MemorystoreHost,
			cfg.MemorystorePort,
			cfg.MemorystoreIndex)
	}
}
