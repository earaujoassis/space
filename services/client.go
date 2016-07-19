package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func CreateNewClient(name, description, scopes, redirectURI string) models.Client {
    var client models.Client = models.Client{
        Name: name,
        Description: description,
        Scopes: scopes,
        RedirectURI: redirectURI,
        Type: models.ConfidentialClient,
    }

    dataStoreSession := datastore.GetDataStoreConnection()
    client.BeforeCreate(dataStoreSession.NewScope(&client))
    dataStoreSession.Create(&client)
    return client
}

func FindOrCreateClient(name string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("name = ?", name).First(&client)
    if dataStoreSession.NewRecord(client) {
        client = models.Client{
            Name: name,
            RedirectURI: "/",
            Scopes: models.PublicScope,
            Type: models.PublicClient,
        }
        client.BeforeCreate(dataStoreSession.NewScope(&client))
        dataStoreSession.Create(&client)
    }
    return client
}

func FindClientByKey(key string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("key = ?", key).First(&client)
    return client
}

func ClientAuthentication(key, secret string) models.Client {
    var client models.Client

    client = FindClientByKey(key)
    if client.ID != 0 && client.Secret == secret {
        return client
    }
    return models.Client{}
}
