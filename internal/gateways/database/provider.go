package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

type DatabaseProvider interface {
	Connect(cfg *config.Config) (*gorm.DB, error)
}

type PostgresProvider struct{}

func (p *PostgresProvider) Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseDSN()
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

type SQLiteProvider struct{}

func (s *SQLiteProvider) Connect(cfg *config.Config) (*gorm.DB, error) {
	filePath := cfg.DatabaseFilepath()
	return gorm.Open(sqlite.Open(filePath), &gorm.Config{})
}

func NewDatabaseProvider(cfg *config.Config) DatabaseProvider {
	switch cfg.Environment {
	case config.Production, config.Development, config.Integration:
		return &PostgresProvider{}
	case config.Test:
		return &SQLiteProvider{}
	default:
		logs.Propagate(logs.LevelPanic, "gateway misconfigured")
		return nil
	}
}
