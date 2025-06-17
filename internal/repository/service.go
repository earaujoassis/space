package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
)

type ServiceRepository struct {
	*BaseRepository[models.Service]
}

func NewServiceRepository(db *database.DatabaseService) *ServiceRepository {
	return &ServiceRepository{
		BaseRepository: NewBaseRepository[models.Service](db),
	}
}

// Create creates a new service application entry
func (r *ServiceRepository) Create(name, description, canonicalURI, logoURI string) models.Service {
	var service models.Service = models.Service{
		Name:         name,
		Description:  description,
		CanonicalURI: canonicalURI,
		LogoURI:      logoURI,
		Type:         models.PublicService,
	}

	r.db.GetDB().Create(&service)

	return service
}

// GetAll lists all services applications
func (r *ServiceRepository) GetAll() []models.Service {
	var services []models.Service

	r.db.GetDB().
		Raw("SELECT services.uuid, services.name, services.description, services.canonical_uri, services.logo_uri " +
			"FROM services " +
			"ORDER BY services.name ASC").
		Scan(&services)

	return services
}
