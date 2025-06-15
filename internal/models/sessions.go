package models

import (
	"time"

	"gorm.io/gorm"
)

// Session model/struct
type Session struct {
	Model
	UUID        string `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
	User        User   `gorm:"not null" validate:"required" json:"-"`
	UserID      uint   `gorm:"not null" json:"-"`
	Client      Client `gorm:"not null" validate:"required" json:"-"`
	ClientID    uint   `gorm:"not null" json:"-"`
	Moment      int64  `gorm:"not null" json:"moment"`
	ExpiresIn   int64  `gorm:"not null;default:0" json:"expires_in"`
	IP          string `gorm:"not null;index" validate:"required" json:"-"`
	UserAgent   string `gorm:"not null" validate:"required" json:"-"`
	Invalidated bool   `gorm:"not null;default:false"`
	Token       string `gorm:"not null;unique;index" validate:"omitempty,alphanum|jwt" json:"token"`
	TokenType   string `gorm:"not null;index" validate:"required,token" json:"token_type"`
	Scopes      string `gorm:"not null" validate:"required,scope" json:"-"`
}

func expirationLengthForTokenType(tokenType string) int64 {
	switch tokenType {
	case AccessToken:
		return largestExpirationLength
	case RefreshToken:
		return refreshableExpirationLength
	case GrantToken:
		return machineryExpirationLength
	default:
		return defaultExpirationLength
	}
}

// BeforeSave Session model/struct hook
func (session *Session) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", session)
}

// BeforeCreate Session model/struct hook
func (session *Session) BeforeCreate(tx *gorm.DB) error {
	if session.Token == "" {
		session.Token = GenerateRandomString(64)
	}
	session.UUID = generateUUID()
	session.Moment = time.Now().UTC().Unix()
	session.ExpiresIn = expirationLengthForTokenType(session.TokenType)
	return nil
}

// WithinExpirationWindow checks if a Session entry is still valid (time-based)
func (session *Session) WithinExpirationWindow() bool {
	now := time.Now().UTC().Unix()
	return now <= session.Moment+session.ExpiresIn
}

// GrantsReadAbility checks if a session entry has read-ability
func (session *Session) GrantsReadAbility() bool {
	return session.Scopes == ReadScope || session.Scopes == WriteScope || session.Scopes == OpenIDScope
}

// GrantsWriteAbility checks if a session entry has write-ability
func (session *Session) GrantsWriteAbility() bool {
	return session.Scopes == WriteScope
}

func HasValidScopes(requestedScopes []string) bool {
	validScopes := []string{PublicScope, ReadScope, OpenIDScope, ProfileScope}
	validSet := make(map[string]bool)
	for _, scope := range validScopes {
		validSet[scope] = true
	}

	for _, requested := range requestedScopes {
		if !validSet[requested] {
			return false
		}
	}

	return true
}
