package workers

import (
	"context"

	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/redis"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/repository"
)

const (
	TypeTokensCleanup = "tokens:cleanup"
)

type TokenCleanupProcessor struct {
	repositories *repository.RepositoryManager
}

func NewTokenCleanupProcessor(cfg *config.Config) *TokenCleanupProcessor {
	db, err := database.NewDatabaseService(cfg)
	if err != nil {
		logs.Propagate(logs.Panic, err.Error())
	}
	ms, err := redis.NewMemoryService(cfg)
	if err != nil {
		logs.Propagate(logs.Panic, err.Error())
	}
	manager := repository.NewRepositoryManager(db, ms)
	return &TokenCleanupProcessor{
		repositories: manager,
	}
}

func (p *TokenCleanupProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	err := p.repositories.Sessions().InvalidateStaleSessions()
	return err
}
