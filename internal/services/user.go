package services

import (
	"github.com/earaujoassis/space/internal/datastore"
	"github.com/earaujoassis/space/internal/models"
)

func SaveUser(user *models.User) {
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Save(user)
}

func CreateNewUser(user *models.User) (bool, error) {
	datastoreSession := datastore.GetDatastoreConnection()
	if user.Client.IsNewRecord() {
		user.Client = FindOrCreateClient(DefaultClient)
	}
	if user.Language.IsNewRecord() {
		user.Language = FindOrCreateLanguage("English", "en-US")
	}
	result := datastoreSession.Create(user)
	return result.RowsAffected >= 1, result.Error
}

// FindUserByAccountHolder gets an user by its account holder (username or email)
func FindUserByAccountHolder(holder string) models.User {
	var user models.User
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Preload("Client").Preload("Language").Where("username = ? OR email = ?", holder, holder).First(&user)
	return user
}

// FindUserByPublicID gets an user by its public ID (used by client applications)
func FindUserByPublicID(publicID string) models.User {
	var user models.User
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Preload("Client").Preload("Language").Where("public_id = ?", publicID).First(&user)
	return user
}

// FindUserByUUID gets an user by its UUID (internal use only)
func FindUserByUUID(uuid string) models.User {
	var user models.User
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Preload("Client").Preload("Language").Where("uuid = ?", uuid).First(&user)
	return user
}

// FindUserByID gets an user by its ID (internal use only)
func FindUserByID(id uint) models.User {
	var user models.User
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Preload("Client").Preload("Language").Where("id = ?", id).First(&user)
	return user
}

// ActiveClientsForUser lists client applications for a given user
func ActiveClientsForUser(userIID uint) []models.Client {
	var clients []models.Client

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Raw("SELECT DISTINCT clients.uuid, clients.name, clients.description, clients.canonical_uri "+
			"FROM clients JOIN sessions ON clients.id = sessions.client_id "+
			"WHERE sessions.token_type IN ('access_token', 'refresh_token') AND sessions.invalidated = false AND "+
			"sessions.user_id = ?;", userIID).
		Scan(&clients)
	return clients
}
