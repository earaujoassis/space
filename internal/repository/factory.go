package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/memory"
)

type RepositoryFactory struct {
	db *database.DatabaseService
	ms *memory.MemoryService
}

func NewRepositoryFactory(db *database.DatabaseService, ms *memory.MemoryService) *RepositoryFactory {
	return &RepositoryFactory{
		db: db,
		ms: ms,
	}
}

func (f *RepositoryFactory) NewActionRepository() *ActionRepository {
	return NewActionRepository(f.ms)
}

func (f *RepositoryFactory) NewClientRepository() *ClientRepository {
	return NewClientRepository(f.db)
}

func (f *RepositoryFactory) NewLanguageRepository() *LanguageRepository {
	return NewLanguageRepository(f.db)
}

func (f *RepositoryFactory) NewNonceRepository() *NonceRepository {
	return NewNonceRepository(f.ms)
}

func (f *RepositoryFactory) NewServiceRepository() *ServiceRepository {
	return NewServiceRepository(f.db)
}

func (f *RepositoryFactory) NewSessionRepository() *SessionRepository {
	return NewSessionRepository(f.db)
}

func (f *RepositoryFactory) NewUserRepository() *UserRepository {
	return NewUserRepository(f.db)
}

func (f *RepositoryFactory) NewEmailRepository() *EmailRepository {
	return NewEmailRepository(f.db)
}

func (f *RepositoryFactory) NewSettingRepository() *SettingRepository {
	return NewSettingRepository(f.db)
}
