package factory

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/redis"
	"github.com/earaujoassis/space/internal/repository"
)

type TestRepositoryFactory struct {
	manager *repository.RepositoryManager
}

func NewTestRepositoryFactory(db *database.DatabaseService, ms *redis.MemoryService) *TestRepositoryFactory {
	return &TestRepositoryFactory{
		manager: repository.NewRepositoryManager(db, ms),
	}
}
