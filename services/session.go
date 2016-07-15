package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func FindSessionByUUID(uuid string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("User").Where("uuid = ? AND invalidated = false", uuid).First(&session)
    return session
}

func FindSessionByToken(token, tokenType string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("User").Where("token = ? AND token_type = ?", token, tokenType).First(&session)
    return session
}

func InvalidateSession(session models.Session) {
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Model(&session).Select("invalidated").Update("invalidated", true)
}
