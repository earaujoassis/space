package main

func FindOrCreateClient(name string) Client {
    var client Client
    dataStoreSession := GetDataStoreConnection()
    dataStoreSession.Where("name = ?", name).First(&client)
    if dataStoreSession.NewRecord(client) {
        client = Client{Name: name}
        client.BeforeCreate(dataStoreSession.NewScope(&client))
        dataStoreSession.Create(&client)
    }
    return client
}

func FindOrCreateLanguage(name string, isoCode string) Language {
    var language Language
    dataStoreSession := GetDataStoreConnection()
    dataStoreSession.Where("name = ? AND iso_code = ?", name, isoCode).First(&language)
    if dataStoreSession.NewRecord(language) {
        language = Language{Name: name, IsoCode: isoCode}
        dataStoreSession.Create(&language)
    }
    return language
}
