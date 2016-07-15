package datastore

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/config"
)

var dataStore *gorm.DB

func init() {
    GetDataStoreConnection().AutoMigrate(&models.Client{},
        &models.Language{},
        &models.User{},
        &models.Session{})
}

func GetDataStoreConnection() *gorm.DB {
    if dataStore != nil {
        return dataStore
    }
    var err error
    var databaseName = fmt.Sprintf("%v_%v",
        config.GetConfig("datastore.database_prefix"),
        config.GetConfig("environment"))
    var databaseConnectionData = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
        config.GetConfig("datastore.host"),
        config.GetConfig("datastore.user"),
        databaseName,
        config.GetConfig("datastore.sslmode"),
        config.GetConfig("datastore.password"),
    )
    fmt.Printf("Connected to the following data store: %v\n", databaseConnectionData)
    dataStore, err = gorm.Open("postgres", databaseConnectionData)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
    }
    return dataStore
}
