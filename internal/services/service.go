package services

import (
	"github.com/earaujoassis/space/internal/datastore"
	"github.com/earaujoassis/space/internal/models"
)

// CreateNewService creates a new service application entry
func CreateNewService(name, description, canonicalURI, logoURI string) *models.Service {
	var service models.Service = models.Service{
		Name:         name,
		Description:  description,
		CanonicalURI: canonicalURI,
		LogoURI:      logoURI,
		Type:         models.PublicService,
	}

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Create(&service)
	return &service
}

// Services lists all services applications
func Services() []models.Service {
	var services []models.Service

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Raw("SELECT services.uuid, services.name, services.description, services.canonical_uri, services.logo_uri " +
			"FROM services " +
			"ORDER BY services.name ASC").
		Scan(&services)
	return services
}
