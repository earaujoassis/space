package services

import (
	"github.com/earaujoassis/space/internal/datastore"
	"github.com/earaujoassis/space/internal/models"
)

const (
	// DefaultClient is the default (and internal) client application
	DefaultClient = "Jupiter"
)

func SaveClient(client *models.Client) {
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Save(client)
}

// CreateNewClient creates a new client application entry
func CreateNewClient(client *models.Client) (bool, error) {
	datastoreSession := datastore.GetDatastoreConnection()
	result := datastoreSession.Create(&client)
	return result.RowsAffected >= 1, result.Error
}

// FindOrCreateClient attempts to find a client application by its name; otherwise, it creates a new one
func FindOrCreateClient(name string) models.Client {
	var client models.Client

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Where("name = ?", name).First(&client)
	if client.IsNewRecord() {
		client = models.Client{
			Name:         name,
			Secret:       models.GenerateRandomString(64),
			CanonicalURI: []string{"localhost"},
			RedirectURI:  []string{"/"},
			Scopes:       models.PublicScope,
			Type:         models.PublicClient,
		}
		datastoreSession.Create(&client)
	}
	return client
}

// FindClientByKey gets a client application by its key
func FindClientByKey(key string) models.Client {
	var client models.Client

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Where("key = ?", key).First(&client)
	return client
}

// FindClientByUUID gets a client application by its UUID
func FindClientByUUID(uuid string) models.Client {
	var client models.Client

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Where("uuid = ?", uuid).First(&client)
	return client
}

// FindClientByID gets a client application by its ID
func FindClientByID(id uint) models.Client {
	var client models.Client

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Where("id = ?", id).First(&client)
	return client
}

// ClientAuthentication gets a client application by its key-secret pair
func ClientAuthentication(key, secret string) models.Client {
	client := FindClientByKey(key)
	if client.ID != 0 && client.Authentic(secret) {
		return client
	}
	return models.Client{}
}

// ActiveClients lists all client applications
func ActiveClients() []models.Client {
	var clients []models.Client

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Raw("SELECT clients.uuid, clients.name, clients.description, clients.scopes, clients.canonical_uri, clients.redirect_uri " +
			"FROM clients " +
			"WHERE clients.name != 'Jupiter' ORDER BY clients.created_at ASC").
		Scan(&clients)
	return clients
}
