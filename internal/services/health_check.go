package services

import (
    "github.com/earaujoassis/space/internal/datastore"
)

func IsDatastoreConnectedAndHealthy() bool {
    var count struct{
        Count int64
    }

    datastoreSession := datastore.GetDatastoreConnection()
    datastoreSession.
        Raw("SELECT count(*) AS count FROM clients;").
        Scan(&count)

    return count.Count >= 0
}
