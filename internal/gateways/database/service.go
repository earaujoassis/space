package database

import (
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/config"
)

type DatabaseService struct {
	db       *gorm.DB
	provider DatabaseProvider
	config   *config.Config
}

func NewDatabaseService(cfg *config.Config) (*DatabaseService, error) {
	provider := NewDatabaseProvider(cfg)

	db, err := provider.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &DatabaseService{
		db:       db,
		provider: provider,
		config:   cfg,
	}, nil
}

func (ds *DatabaseService) GetDB() *gorm.DB {
	return ds.db
}

func (ds *DatabaseService) GetConfig() *config.Config {
	return ds.config
}

func (ds *DatabaseService) GetStorageSecret() []byte {
	return []byte(ds.config.StorageSecret)
}

func (ds *DatabaseService) Close() error {
	sqlDB, err := ds.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
