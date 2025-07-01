package tasks

import (
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/workers"
)

func Scheduler(cfg *config.Config) {
	redisAddr := fmt.Sprintf("%s:%d", cfg.MemorystoreHost, cfg.MemorystorePort)
	scheduler := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: redisAddr, DB: cfg.MemorystoreIndex},
		&asynq.SchedulerOpts{Location: time.UTC},
	)

	if entryID, err := scheduler.Register("@every 15m", asynq.NewTask(workers.TypeTokensCleanup, nil)); err != nil {
		logs.Propagatef(logs.Error, "could not register schedule: %v", err)
	} else {
		log.Printf("Schedule registered with ID: %s\n", entryID)
	}

	if err := scheduler.Run(); err != nil {
		logs.Propagatef(logs.Error, "could not run scheduler: %v", err)
	}
}
