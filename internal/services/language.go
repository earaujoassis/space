package services

import (
    "github.com/earaujoassis/space/internal/datastore"
    "github.com/earaujoassis/space/internal/models"
)

// FindOrCreateLanguage attempts to obtain a Language entry through its `name` and ISO Code;
//      if that's not available, it creates one
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
