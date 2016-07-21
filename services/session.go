package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func CreateSession(user models.User, client models.Client, ip, userAgent, scopes, tokenType string) models.Session {
    var session models.Session = models.Session{
        User: user,
        Client: client,
        Ip: ip,
        UserAgent: userAgent,
        Scopes: scopes,
        TokenType: tokenType,
    }
    dataStore := datastore.GetDataStoreConnection()
    dataStore.Create(&session)
    return session
}

func FindSessionByUUID(uuid string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Preload("Client").
        Preload("User").
        Preload("User.Client").
        Preload("User.Language").
        Where("uuid = ? AND invalidated = false", uuid).
        First(&session)
    return session
}

func FindSessionByToken(token, tokenType string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Preload("Client").
        Preload("User").
        Preload("User.Client").
        Preload("User.Language").
        Where("token = ? AND token_type = ? AND invalidated = false", token, tokenType).
        First(&session)
    return session
}

func InvalidateSession(session models.Session) {
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Model(&session).Select("invalidated").Update("invalidated", true)
}

func ActiveSessionsForClient(clientKey string) int64 {
    var count struct{
        Count int64
    }

    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.
        Raw("SELECT count(*) AS count FROM sessions JOIN clients ON clients.id = sessions.client_id WHERE clients.key = ? AND invalidated = false AND token_type = ? OR token_type = ?;",
            clientKey, models.AccessToken, models.RefreshToken).
        Scan(&count)
    return count.Count
}

func SessionGrantsReadAbility(session models.Session) bool {
    return session.Scopes == models.ReadScope || session.Scopes == models.ReadWriteScope
}

func SessionGrantsWriteAbility(session models.Session) bool {
    return session.Scopes == models.ReadWriteScope
}
