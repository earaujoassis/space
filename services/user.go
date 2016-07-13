package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func FindUserByAccountHolder(holder string) models.User {
    var user models.User
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Preload("Client").Preload("Language").Where("username = ? OR email = ?", holder, holder).First(&user)
    if dataStoreSession.NewRecord(user) {
        return models.User{}
    } else {
        return user
    }
}
