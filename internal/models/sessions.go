package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

const (
	// AccessToken token type
	AccessToken string = "access_token"
	// RefreshToken token type
	RefreshToken string = "refresh_token"
	// GrantToken token type
	GrantToken string = "grant_token"

	// PublicScope session scope
	// This is used by public clients (they can't read or write user data)
	PublicScope string = "public"
	// ReadScope session scope
	// This is used by confidential clients (they can only read user data)
	ReadScope string = "read"
	// WriteScope session scope
	// No client is allowed to hold this scope (they can't write user data)
	WriteScope string = "write"
	// OpenIDScope session scope
	// This is used for OpenID Connect
	OpenIDScope string = "openid"
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
	Token       string `gorm:"not null;unique;index" validate:"omitempty,alphanum" json:"token"`
	TokenType   string `gorm:"not null;index" validate:"required,token" json:"token_type"`
	Scopes      string `gorm:"not null" validate:"required,scope" json:"-"`
}

func validScope(fl validator.FieldLevel) bool {
	scope := fl.Field().String()
	if scope != PublicScope && scope != ReadScope && scope != WriteScope && scope != OpenIDScope {
		return false
	}
	return true
}

func validTokenType(fl validator.FieldLevel) bool {
	tokenType := fl.Field().String()
	if tokenType != AccessToken && tokenType != RefreshToken && tokenType != GrantToken {
		return false
	}
	return true
}

func expirationLengthForTokenType(tokenType string) int64 {
	switch tokenType {
	case AccessToken:
		return largestExpirationLength
	case RefreshToken:
		return eternalExpirationLength
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
	session.Token = GenerateRandomString(64)
	session.UUID = generateUUID()
	session.Moment = time.Now().UTC().Unix()
	session.ExpiresIn = expirationLengthForTokenType(session.TokenType)
	return nil
}

// WithinExpirationWindow checks if a Session entry is still valid (time-based)
func (session *Session) WithinExpirationWindow() bool {
	now := time.Now().UTC().Unix()
	return session.ExpiresIn == eternalExpirationLength || session.Moment+session.ExpiresIn >= now
}
