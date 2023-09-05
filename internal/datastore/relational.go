package datastore

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/config"
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
	log.Printf("Connected to the following data store: %v\n", safeDatabaseConnectionData)
	datastore, err = gorm.Open(postgres.Open(databaseConnectionData), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
	}
}

// GetDatastoreConnection is used to obtain a connection with
//
//	the Postgres datastore
func GetDatastoreConnection() *gorm.DB {
	if datastore != nil {
		return datastore
	}

	InitConnection()

	return datastore
}
