package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
)

type ClientRepository struct {
	*BaseRepository[models.Client]
}

func NewClientRepository(db *database.DatabaseService) *ClientRepository {
	return &ClientRepository{
		BaseRepository: NewBaseRepository[models.Client](db),
	}
}

// FindOrCreate attempts to find a client application by its name;
// otherwise, it creates a new one
func (r *ClientRepository) FindOrCreate(name string) models.Client {
	var client models.Client

	r.db.GetDB().Where("name = ?", name).First(&client)
	if client.IsNewRecord() {
		client = models.Client{
			Name:         name,
			Secret:       models.GenerateRandomString(64),
			CanonicalURI: []string{"localhost"},
			RedirectURI:  []string{"/"},
			Scopes:       models.PublicScope,
			Type:         models.PublicClient,
		}
		r.db.GetDB().Create(&client)
	}

	return client
}

// FindByKey gets a client application by its key
func (r *ClientRepository) FindByKey(key string) models.Client {
	var client models.Client

	r.db.GetDB().Where("key = ?", key).First(&client)

	return client
}

// FindByUUID gets a client application by its UUID
func (r *ClientRepository) FindByUUID(uuid string) models.Client {
	var client models.Client

	r.db.GetDB().Where("uuid = ?", uuid).First(&client)

	return client
}

// Authentication gets a client application by its key-secret pair
func (r *ClientRepository) Authentication(key, secret string) models.Client {
	client := r.FindByKey(key)
	if client.IsSavedRecord() && client.Authentic(secret) {
		return client
	}

	return models.Client{}
}

// GetActive lists all active client applications
func (r *ClientRepository) GetActive() []models.Client {
	clients := make([]models.Client, 0)

	r.db.GetDB().
		Raw(`SELECT clients.uuid, clients.name, clients.description, clients.scopes, clients.canonical_uri, clients.redirect_uri
			FROM clients
			WHERE clients.name != 'Jupiter' ORDER BY clients.created_at ASC`).
		Scan(&clients)

	return clients
}
