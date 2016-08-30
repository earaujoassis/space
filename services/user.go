package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func FindUserByAccountHolder(holder string) models.User {
    var user models.User
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("Language").Where("username = ? OR email = ?", holder, holder).First(&user)
    return user
}

func FindUserByPublicId(publicId string) models.User {
    var user models.User
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("Language").Where("public_id = ?", publicId).First(&user)
    return user
}

func FindUserByUUID(uuid string) models.User {
    var user models.User
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("Language").Where("uuid = ?", uuid).First(&user)
    return user
}
