package tasks

import (
	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/logs/plugins"
	"github.com/earaujoassis/space/internal/workers"
)

func Workers(cfg *config.Config) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: cfg.MemoryDNS(),
			DB:   cfg.MemorystoreIndex,
		},
		asynq.Config{
			Concurrency: 5,
			Queues: map[string]int{
				"critical": 3,
				"default":  2,
				"low":      1,
			},
			ErrorHandler: plugins.SentryErrorHandler(),
		},
	)

	mux := asynq.NewServeMux()
	mux.Use(plugins.SentryMiddleware())
	mux.Handle(workers.TypeTokensCleanup, workers.NewTokenCleanupProcessor(cfg))

	if err := srv.Run(mux); err != nil {
		logs.Propagatef(logs.LevelError, "could not run worker server: %v", err)
	}
}
