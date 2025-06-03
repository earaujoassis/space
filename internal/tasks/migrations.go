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

func buildDatabaseUrl() string {
	cfg := config.GetGlobalConfig()
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s_%s?sslmode=%s",
		cfg.DatastoreUser,
		cfg.DatastorePassword,
		cfg.DatastoreHost,
		cfg.DatastorePort,
		cfg.DatastoreNamePrefix,
		config.Environment(),
		cfg.DatastoreSslMode)
}

func createMigrator(relativePath string) (*migrate.Migrate, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	migrationsPath := fmt.Sprintf("file:///%s", filepath.Join(pwd, relativePath))
	db, err := sql.Open("postgres", buildDatabaseUrl())
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

func RunMigrations(relativePath string) {
	migrator, err := createMigrator(relativePath)
	if err != nil {
		logs.Propagate(logs.Panic, err.Error())
	}
	if err = migrator.Up(); err != nil && err != migrate.ErrNoChange {
		logs.Propagate(logs.Panic, err.Error())
	}
}

func RollbackMigrations(relativePath string) {
	migrator, err := createMigrator(relativePath)
	if err != nil {
		logs.Propagate(logs.Panic, err.Error())
	}
	if err = migrator.Down(); err != nil {
		logs.Propagate(logs.Panic, err.Error())
	}
}
