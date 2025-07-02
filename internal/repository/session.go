package repository

import (
	"time"

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
	if session.IsSavedRecord() {
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
	if session.IsSavedRecord() {
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

func (r *SessionRepository) ApplicationSessions(user models.User) []models.Session {
	sessions := make([]models.Session, 0)

	r.db.GetDB().
		Raw(`SELECT sessions.uuid, sessions.ip, sessions.user_agent
			FROM sessions
			WHERE token_type = 'application_token' AND invalidated = false AND user_id = ?
			ORDER BY sessions.created_at DESC;`, user.ID).
		Scan(&sessions)

	return sessions
}

func (r *SessionRepository) ApplicationSessionsWithActive(user models.User, activeSession models.Session) []models.Session {
	sessions := r.ApplicationSessions(user)

	for i := range sessions {
		sessions[i].Current = (sessions[i].UUID == activeSession.UUID)
	}

	return sessions
}

// ActiveForClient gets the number of active sessions for a given user in a client application
func (r *SessionRepository) ActiveForClient(client models.Client, user models.User) int64 {
	var count struct {
		Count int64
	}

	r.db.GetDB().
		Raw(`SELECT count(*) AS count
			FROM sessions WHERE token_type IN ('access_token', 'refresh_token') AND
				invalidated = false AND
				client_id = ? AND
				user_id = ?;`, client.ID, user.ID).
		Scan(&count)
	return count.Count
}

// RevokeAccess revokes client application access to user's data
func (r *SessionRepository) RevokeAccess(client models.Client, user models.User) {
	now := time.Now().UTC()
	r.db.GetDB().
		Exec(`UPDATE sessions SET invalidated = true, updated_at = ?
			WHERE token_type IN ('access_token', 'refresh_token') AND
				invalidated = false AND
				client_id = ? AND
				user_id = ?;`, now, client.ID, user.ID)
}

func (r *SessionRepository) InvalidateStaleSessions() error {
	now := time.Now().UTC()
	nowUnix := time.Now().UTC().Unix()
	return r.db.GetDB().
		Exec(`UPDATE sessions SET invalidated = true, updated_at = ?
			WHERE invalidated = false AND
				(moment + expires_in) < ?;`, now, nowUnix).Error
}
