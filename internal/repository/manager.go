package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/memory"
)

type RepositoryManager struct {
	factory *RepositoryFactory
}

func NewRepositoryManager(db *database.DatabaseService, ms *memory.MemoryService) *RepositoryManager {
	return &RepositoryManager{
		factory: NewRepositoryFactory(db, ms),
	}
}

func (rm *RepositoryManager) Actions() *ActionRepository {
	return rm.factory.NewActionRepository()
}

func (rm *RepositoryManager) Clients() *ClientRepository {
	return rm.factory.NewClientRepository()
}

func (rm *RepositoryManager) Languages() *LanguageRepository {
	return rm.factory.NewLanguageRepository()
}

func (rm *RepositoryManager) Nonces() *NonceRepository {
	return rm.factory.NewNonceRepository()
}

func (rm *RepositoryManager) Services() *ServiceRepository {
	return rm.factory.NewServiceRepository()
}

func (rm *RepositoryManager) Sessions() *SessionRepository {
	return rm.factory.NewSessionRepository()
}

func (rm *RepositoryManager) Users() *UserRepository {
	return rm.factory.NewUserRepository()
}
