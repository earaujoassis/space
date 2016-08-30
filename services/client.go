package services

import (
    "fmt"
    "net/url"

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

func FindClientByUUID(uuid string) models.Client {
    var client models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("uuid = ?", uuid).First(&client)
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

func ActiveClientsForUser(internalUserId uint) []models.Client {
    var clients []models.Client

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Raw("SELECT DISTINCT clients.uuid, clients.name, clients.description, clients.redirect_uri " +
            "FROM clients JOIN sessions ON clients.id = sessions.client_id " +
            "WHERE sessions.user_id = ? AND sessions.invalidated = false AND sessions.token_type = ? OR sessions.token_type = ?;",
            internalUserId, models.AccessToken, models.RefreshToken).
        Scan(&clients)
    for i := range clients {
        client := clients[i]
        u, _ := url.Parse(client.RedirectURI)
        client.RedirectURI = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
        clients[i] = client
    }
    return clients
}
