package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

const (
    // AccessToken token type
    AccessToken               string = "access_token"
    // RefreshToken token type
    RefreshToken              string = "refresh_token"
    // GrantToken token type
    GrantToken                string = "grant_token"

    // PublicScope session scope
    // This is used by public clients (they can't read or write user data)
    PublicScope               string = "public"
    // ReadScope session scope
    // This is used by confidential clients (they can only read user data)
    ReadScope                 string = "read"
    // ReadWriteScope session scope
    // No client is allowed to hold this scope (they can't write user data)
    ReadWriteScope            string = "read_write"
)

// Session model/struct
type Session struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    User User                   `gorm:"not null" validate:"exists" json:"-"`
    UserID uint                 `gorm:"not null" json:"-"`
    Client Client               `gorm:"not null" validate:"exists" json:"-"`
    ClientID uint               `gorm:"not null" json:"-"`
    Moment int64                `gorm:"not null" json:"moment"`
    ExpiresIn int64             `gorm:"not null;default:0" json:"expires_in"`
    IP string                   `gorm:"not null;index" validate:"required" json:"-"`
    UserAgent string            `gorm:"not null" validate:"required" json:"-"`
    Invalidated bool            `gorm:"not null;default:false"`
    Token string                `gorm:"not null;unique;index" validate:"omitempty,alphanum" json:"token"`
    TokenType string            `gorm:"not null;index" validate:"required,token" json:"token_type"`
    Scopes string               `gorm:"not null" validate:"required,scope" json:"-"`
}

func validScope(top interface{}, current interface{}, field interface{}, param string) bool {
    scope := field.(string)
    if scope != PublicScope && scope != ReadScope && scope != ReadWriteScope {
        return false
    }
    return true
}

func validTokenType(top interface{}, current interface{}, field interface{}, param string) bool {
    tokenType := field.(string)
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
func (session *Session) BeforeSave(scope *gorm.Scope) error {
    return validateModel("validate", session)
}

// BeforeCreate Session model/struct hook
func (session *Session) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("Token", GenerateRandomString(64))
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("Moment", time.Now().UTC().Unix())
    scope.SetColumn("ExpiresIn", expirationLengthForTokenType(session.TokenType))
    return nil
}

// WithinExpirationWindow checks if a Session entry is still valid (time-based)
func (session *Session) WithinExpirationWindow() bool {
    now := time.Now().UTC().Unix()
    return session.ExpiresIn == eternalExpirationLength || session.Moment + session.ExpiresIn >= now
}
