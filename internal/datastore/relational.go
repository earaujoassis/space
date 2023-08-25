package datastore

import (
    "fmt"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"

    "github.com/earaujoassis/space/internal/config"
    "github.com/earaujoassis/space/internal/models"
)

var datastore *gorm.DB

// Start is used to setup the models within the application
func Start() {
    GetDatastoreConnection().AutoMigrate(&models.Client{},
        &models.Language{},
        &models.User{},
        &models.Session{})
}

// GetDatastoreConnection is used to obtain a connection with
//      the Postgres datastore
func GetDatastoreConnection() *gorm.DB {
    if datastore != nil {
        return datastore
    }
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
    fmt.Printf("Connected to the following data store: %v\n", safeDatabaseConnectionData)
    datastore, err = gorm.Open(postgres.Open(databaseConnectionData), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
    }
    return datastore
}
