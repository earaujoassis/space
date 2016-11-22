package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func CreateNewClient(name, description, secret, scopes, canonicalURI, redirectURI string) models.Client {
    var client models.Client = models.Client{
        Name: name,
        Description: description,
        Secret: secret,
        Scopes: scopes,
        CanonicalURI: canonicalURI,
        RedirectURI: redirectURI,
        Type: models.ConfidentialClient,
    }

    dataStoreSession := datastore.GetDataStoreConnection()
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
            Secret: models.GenerateRandomString(64),
            CanonicalURI: "localhost",
            RedirectURI: "/",
            Scopes: models.PublicScope,
            Type: models.PublicClient,
        }
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

func FindClientByUUID(uuid string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("uuid = ?", uuid).First(&client)
    return client
}

func ClientAuthentication(key, secret string) models.Client {
    var client models.Client

    client = FindClientByKey(key)
    if client.ID != 0 && client.Authentic(secret) {
        return client
    }
    return models.Client{}
}

func ActiveClientsForUser(userIID uint) []models.Client {
    var clients []models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Raw("SELECT DISTINCT clients.uuid, clients.name, clients.description, clients.canonical_uri " +
            "FROM clients JOIN sessions ON clients.id = sessions.client_id " +
            "WHERE sessions.token_type IN ('access_token', 'refresh_token') AND sessions.invalidated = false AND " +
            "sessions.user_id = ?;", userIID).
        Scan(&clients)
    return clients
}
