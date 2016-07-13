package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func FindSessionByUUID(uuid string) models.Session {
    var session models.Session
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("User").Where("uuid = ? AND invalidated = false", uuid).First(&session)
    if dataStoreSession.NewRecord(session) {
        return models.Session{}
    } else {
        return session
    }
}
