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

// GetAll lists all services applications
func (r *ServiceRepository) GetAll() []models.Service {
	services := make([]models.Service, 0)

	r.db.GetDB().
		Raw(`SELECT services.uuid, services.name, services.description, services.canonical_uri, services.logo_uri
			FROM services
			ORDER BY services.name ASC`).
		Scan(&services)

	return services
}
