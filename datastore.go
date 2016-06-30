package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDataStoreConnection() *gorm.DB {
    var databaseName = fmt.Sprintf("%v_%v", GetConfig("datastore.database_prefix"), GetConfig("environment"))
    var databaseConnectionData = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
        GetConfig("datastore.host"),
        GetConfig("datastore.user"),
        databaseName,
        GetConfig("datastore.sslmode"),
        GetConfig("datastore.password"),
    )
    fmt.Printf("Connected to the following data store: %v\n", databaseConnectionData)
    dataStore, err := gorm.Open("postgres", databaseConnectionData)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect datastore: %v\n", err))
    }
    return dataStore
}
