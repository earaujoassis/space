package services

import (
	datastore "github.com/earaujoassis/space/internal/gateways/postgres"
)

func IsDatastoreConnectedAndHealthy() bool {
	var count struct {
		Count int64
	}

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Raw("SELECT count(*) AS count FROM clients;").
		Scan(&count)

	return count.Count >= 0
}
