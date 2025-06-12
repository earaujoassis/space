package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
)

var datastore *gorm.DB

// InitConnection is used to start a connection with the Postgres datastore
func InitConnection() {
	var err error
	var cfg config.Config = config.GetGlobalConfig()
	var databaseName = fmt.Sprintf("%v_%v",
		cfg.DatastoreNamePrefix, config.Environment())
	var databaseConnectionData = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		cfg.DatastoreHost,
		cfg.DatastorePort,
		cfg.DatastoreUser,
		databaseName,
		cfg.DatastoreSslMode,
		cfg.DatastorePassword,
	)
	var safeDatabaseConnectionData = fmt.Sprintf("host=%s port=%d user=%s dbname=%s",
		cfg.DatastoreHost,
		cfg.DatastorePort,
		cfg.DatastoreUser,
		databaseName,
	)
	logs.Propagatef(logs.Info, "Connected to the following data store: %v\n", safeDatabaseConnectionData)
	datastore, err = gorm.Open(postgres.Open(databaseConnectionData), &gorm.Config{})
	if err != nil {
		logs.Propagatef(logs.Panic, "Failed to connect datastore: %v\n", err)
	}
}
