package services

import (
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func FindOrCreateLanguage(name, isoCode string) models.Language {
    var language models.Language
    dataStoreSession := datastore.GetDataStoreConnection()
    dataStoreSession.Where("name = ? AND iso_code = ?", name, isoCode).First(&language)
    if dataStoreSession.NewRecord(language) {
        language = models.Language{Name: name, IsoCode: isoCode}
        dataStoreSession.Create(&language)
    }
    return language
}
