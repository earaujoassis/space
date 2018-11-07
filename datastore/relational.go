package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    // Uses Postgres for GORM setup
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/config"
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
    var databaseName = fmt.Sprintf("%v_%v",
        config.GetConfig("SPACE_DATASTORE_NAME_PREFIX"), config.Environment())
    var databaseConnectionData = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
        config.GetConfig("SPACE_DATASTORE_HOST"),
        config.GetConfig("SPACE_DATASTORE_USER"),
        databaseName,
        config.GetConfig("SPACE_DATASTORE_SSL_MODE"),
        config.GetConfig("SPACE_DATASTORE_PASSWORD"),
    )
    fmt.Printf("Connected to the following data store: %v\n", databaseConnectionData)
    dataStore, err = gorm.Open("postgres", databaseConnectionData)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
    }
    return dataStore
}
