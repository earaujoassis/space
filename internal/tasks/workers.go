package tasks

import (
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/workers"
)

func Workers(cfg *config.Config) {
	redisAddr := fmt.Sprintf("%s:%d", cfg.MemorystoreHost, cfg.MemorystorePort)
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr, DB: cfg.MemorystoreIndex},
		asynq.Config{
			Concurrency: 5,
			Queues: map[string]int{
				"critical": 3,
				"default":  2,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.Handle(workers.TypeTokensCleanup, workers.NewTokenCleanupProcessor(cfg))

	if err := srv.Run(mux); err != nil {
		logs.Propagatef(logs.Error, "could not run worker server: %v", err)
	}
}
