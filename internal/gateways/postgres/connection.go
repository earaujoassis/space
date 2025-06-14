package postgres

import (
	"gorm.io/gorm"
)

// GetDatastoreConnection is used to obtain a connection with
//
//	the Postgres datastore
func GetDatastoreConnection() *gorm.DB {
	if datastore != nil {
		return datastore
	}

	InitConnection()

	return datastore
}
