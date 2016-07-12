package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func FindOrCreateClient(name string) models.Client {
    var client models.Client
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("name = ?", name).First(&client)
    if dataStoreSession.NewRecord(client) {
        client = models.Client{Name: name}
        client.BeforeCreate(dataStoreSession.NewScope(&client))
        dataStoreSession.Create(&client)
    }
    return client
}
