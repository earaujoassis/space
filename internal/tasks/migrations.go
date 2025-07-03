package tasks

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

func buildDatabaseUrl(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s_%s?sslmode=%s",
		cfg.DatastoreUser,
		cfg.DatastorePassword,
		cfg.DatastoreHost,
		cfg.DatastorePort,
		cfg.DatastoreNamePrefix,
		cfg.Environment,
		cfg.DatastoreSslMode)
}

func createMigrator(cfg *config.Config, relativePath string) (*migrate.Migrate, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	migrationsPath := fmt.Sprintf("file:///%s", filepath.Join(pwd, relativePath))
	db, err := sql.Open("postgres", buildDatabaseUrl(cfg))
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	migrator, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return migrator, nil
}

func RunMigrations(cfg *config.Config, relativePath string) {
	migrator, err := createMigrator(cfg, relativePath)
	if err != nil {
		logs.Propagate(logs.LevelPanic, err.Error())
	}
	if err = migrator.Up(); err != nil && err != migrate.ErrNoChange {
		logs.Propagate(logs.LevelPanic, err.Error())
	}
}

func RollbackMigrations(cfg *config.Config, relativePath string) {
	migrator, err := createMigrator(cfg, relativePath)
	if err != nil {
		logs.Propagate(logs.LevelPanic, err.Error())
	}
	if err = migrator.Down(); err != nil {
		logs.Propagate(logs.LevelPanic, err.Error())
	}
}
