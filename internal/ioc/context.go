package ioc

import (
	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/internal/notifications"
	"github.com/earaujoassis/space/internal/policy"
	"github.com/earaujoassis/space/internal/repository"
)

type AppContext struct {
	Config       *config.Config
	DB           *database.DatabaseService
	Memory       *memory.MemoryService
	Repositories *repository.RepositoryManager
	FeatureGate  *feature.FeatureGate
	RateLimit    *policy.RateLimitService
	Notifier     *notifications.Notifier
}

func NewAppContext(cfg *config.Config) (*AppContext, error) {
	db, err := database.NewDatabaseService(cfg)
	if err != nil {
		return nil, err
	}
	ms, err := memory.NewMemoryService(cfg)
	if err != nil {
		return nil, err
	}
	rm := repository.NewRepositoryManager(db, ms)
	fg := feature.NewFeatureGate(ms)
	rls := policy.NewRateLimitService(ms)
	ntfr := notifications.NewNotifier(cfg, rm)

	return &AppContext{
		Config:       cfg,
		DB:           db,
		Memory:       ms,
		Repositories: rm,
		FeatureGate:  fg,
		RateLimit:    rls,
		Notifier:     ntfr,
	}, nil
}

func (ctx *AppContext) Close() error {
	ctx.Memory.Close()
	return ctx.DB.Close()
}
