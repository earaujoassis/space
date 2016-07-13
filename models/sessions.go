package models

import (
    "github.com/jinzhu/gorm"
    "gopkg.in/bluesuncorp/validator.v5"
)

type Session struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    User User                   `gorm:"not null" validate:"exists" json:"-"`
    UserID uint                 `gorm:"not null" json:"-"`
    Client Client               `gorm:"not null" validate:"exists" json:"-"`
    ClientID uint               `gorm:"not null" json:"-"`
    Moment int64                `gorm:"not null" validate:"required"`
    Ip string                   `gorm:"not null;index" validate:"required" json:"-"`
    UserAgent string            `gorm:"not null" validate:"required" json:"-"`
    Invalidated bool            `gorm:"default:false"`
    AccessToken string          `gorm:"not null" validate:"required" json:"access_token"`
    RefreshToken string         `gorm:"not null" validate:"required" json:"refresh_token"`
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
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("AccessToken", randStringBytesMaskImprSrc(64))
    scope.SetColumn("RefreshToken", randStringBytesMaskImprSrc(64))
    return nil
}
