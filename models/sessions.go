package models

import (
    "time"

    "github.com/jinzhu/gorm"
    "gopkg.in/bluesuncorp/validator.v5"
)

const (
    AccessToken     string = "access_token"
    RefreshToken    string = "refresh_token"
    GrantToken      string = "grant_token"
    ActionToken     string = "action_token"

    PublicScope     string = "public"
    ReadScope       string = "read"
    ReadWriteScope  string = "read_write"
)

type Session struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    User User                   `gorm:"not null" validate:"exists" json:"-"`
    UserID uint                 `gorm:"not null" json:"-"`
    Client Client               `gorm:"not null" validate:"exists" json:"-"`
    ClientID uint               `gorm:"not null" json:"-"`
    Moment int64                `gorm:"not null"`
    Ip string                   `gorm:"not null;index" validate:"required" json:"-"`
    UserAgent string            `gorm:"not null" validate:"required" json:"-"`
    Invalidated bool            `gorm:"default:false"`
    Token string                `gorm:"not null;unique;index" validate:"omitempty,alphanum" json:"token"`
    TokenType string            `gorm:"not null;index" validate:"required" json:"token_type"`
    Scopes string               `gorm:"not null" json:"-"`
}

func (session *Session) BeforeSave(scope *gorm.Scope) error {
    validate := validator.New("validate", validator.BakedInValidators)
    err := validate.Struct(session)
    if err != nil {
        return err
    }
    return nil
}

func (session *Session) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("Token", randStringBytesMaskImprSrc(64))
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("Moment", time.Now().UTC().Unix())
    return nil
}
