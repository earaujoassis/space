package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    // Uses Postgres for GORM setup
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/earaujoassis/space/internal/config"
    "github.com/earaujoassis/space/internal/models"
)

var dataStore *gorm.DB

// Start is used to setup the models within the application
func Start() {
    GetDataStoreConnection().AutoMigrate(&models.Client{},
        &models.Language{},
        &models.User{},
        &models.Session{})
}

// GetDataStoreConnection is used to obtain a connection with
//      the Postgres datastore
func GetDataStoreConnection() *gorm.DB {
    if dataStore != nil {
        return dataStore
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
    dataStore, err = gorm.Open("postgres", databaseConnectionData)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
    }
    return dataStore
}
