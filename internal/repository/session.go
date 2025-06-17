package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
)

type SessionRepository struct {
	*BaseRepository[models.Session]
}

func NewSessionRepository(db *database.DatabaseService) *SessionRepository {
	return &SessionRepository{
		BaseRepository: NewBaseRepository[models.Session](db),
	}
}

// Create creates a session entry
func (r *SessionRepository) Create(session *models.Session) error {
	return r.db.GetDB().Create(&session).Error
}

// FindByUUID gets a session entry by its UUID
func (r *SessionRepository) FindByUUID(uuid string) models.Session {
	var session models.Session
	r.db.GetDB().
		Preload("Client").
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("uuid = ? AND invalidated = false", uuid).
		First(&session)
	if session.ID != 0 {
		if !session.WithinExpirationWindow() {
			r.Invalidate(&session)
			return models.Session{}
		}
	}
	return session
}

// FindByToken gets a session entry by its token string
func (r *SessionRepository) FindByToken(token, tokenType string) models.Session {
	var session models.Session
	r.db.GetDB().
		Preload("Client").
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("token = ? AND token_type = ? AND invalidated = false", token, tokenType).
		First(&session)
	if session.ID != 0 {
		if !session.WithinExpirationWindow() {
			r.Invalidate(&session)
			return models.Session{}
		}
	}
	return session
}

// Invalidate invalidates a session entry
func (r *SessionRepository) Invalidate(session *models.Session) {
	r.db.GetDB().Model(&session).Select("invalidated").Update("invalidated", true)
}

// ActiveForClient gets the number of active sessions for a given user in a client application
func (r *SessionRepository) ActiveForClient(client models.Client, user models.User) int64 {
	var count struct {
		Count int64
	}

	r.db.GetDB().
		Raw("SELECT count(*) AS count "+
			"FROM sessions WHERE token_type IN ('access_token', 'refresh_token') AND invalidated = false AND "+
			"client_id = ? AND user_id = ?;", client.ID, user.ID).
		Scan(&count)
	return count.Count
}

// RevokeAccess revokes client application access to user's data
func (r *SessionRepository) RevokeAccess(client models.Client, user models.User) {
	r.db.GetDB().
		Exec("UPDATE sessions SET invalidated = true, updated_at = now() "+
			"WHERE token_type IN ('access_token', 'refresh_token') AND invalidated = false AND "+
			"client_id = ? AND user_id = ?;", client.ID, user.ID)
}
