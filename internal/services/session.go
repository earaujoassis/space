package services

import (
	"github.com/earaujoassis/space/internal/datastore"
	"github.com/earaujoassis/space/internal/models"
)

// CreateSession creates a session entry
func CreateSession(user models.User, client models.Client, ip, userAgent, scopes, tokenType string) models.Session {
	var session models.Session = models.Session{
		User:      user,
		Client:    client,
		IP:        ip,
		UserAgent: userAgent,
		Scopes:    scopes,
		TokenType: tokenType,
	}
	datastore := datastore.GetDatastoreConnection()
	result := datastore.Create(&session)
	if count := result.RowsAffected; count > 0 {
		return session
	}
	return models.Session{}
}

// SessionGrantsReadAbility checks if a session entry has read-ability
func SessionGrantsReadAbility(session models.Session) bool {
	return session.Scopes == models.ReadScope || session.Scopes == models.ReadWriteScope
}

// SessionGrantsWriteAbility checks if a session entry has write-ability
func SessionGrantsWriteAbility(session models.Session) bool {
	return session.Scopes == models.ReadWriteScope
}

// FindSessionByUUID gets a session entry by its UUID
func FindSessionByUUID(uuid string) models.Session {
	var session models.Session
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Preload("Client").
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("uuid = ? AND invalidated = false", uuid).
		First(&session)
	if session.ID != 0 {
		if !session.WithinExpirationWindow() {
			InvalidateSession(session)
			return models.Session{}
		}
	}
	return session
}

// FindSessionByToken gets a session entry by its token string
func FindSessionByToken(token, tokenType string) models.Session {
	var session models.Session
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Preload("Client").
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("token = ? AND token_type = ? AND invalidated = false", token, tokenType).
		First(&session)
	if session.ID != 0 {
		if !session.WithinExpirationWindow() {
			InvalidateSession(session)
			return models.Session{}
		}
	}
	return session
}

// InvalidateSession invalidates a session entry
func InvalidateSession(session models.Session) {
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.Model(&session).Select("invalidated").Update("invalidated", true)
}

// ActiveSessionsForClient gets the number of active sessions for a given user in a client application
func ActiveSessionsForClient(clientIID, userIID uint) int64 {
	var count struct {
		Count int64
	}

	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Raw("SELECT count(*) AS count "+
			"FROM sessions WHERE token_type IN ('access_token', 'refresh_token') AND invalidated = false AND "+
			"client_id = ? AND user_id = ?;", clientIID, userIID).
		Scan(&count)
	return count.Count
}

// RevokeClientAccess revokes client application access to user's data
func RevokeClientAccess(clientIID, userIID uint) {
	datastoreSession := datastore.GetDatastoreConnection()
	datastoreSession.
		Exec("UPDATE sessions SET invalidated = true, updated_at = now() "+
			"WHERE token_type IN ('access_token', 'refresh_token') AND invalidated = false AND "+
			"client_id = ? AND user_id = ?;", clientIID, userIID)
}
