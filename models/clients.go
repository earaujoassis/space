package models

import (
    "github.com/jinzhu/gorm"
    "gopkg.in/bluesuncorp/validator.v5"
)

const (
    PublicClient        string = "public"
    ConfidentialClient  string = "confidential"
)

type Client struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    Name string                 `gorm:"not null;unique;index" validate:"required,min=3,max=20" json:"name"`
    Description string          `json:"description"`
    Key string                  `gorm:"not null;unique;index" validate:"required" json:"client_id"`
    Secret string               `gorm:"not null" validate:"required" json:"-"`
    Scopes string               `gorm:"not null" validate:"required" json:"-"`
    RedirectURI string          `gorm:"not null" validate:"required" json:"-"`
    Type string                 `gorm:"not null" validate:"required" json:"client_type"`
}

func (client *Client) BeforeSave(scope *gorm.Scope) error {
    validate := validator.New("validate", validator.BakedInValidators)
    err := validate.Struct(client)
    if err != nil {
        return err
    }
    return nil
}

func (client *Client) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("Secret", randStringBytesMaskImprSrc(64))
    scope.SetColumn("Key", randStringBytesMaskImprSrc(32))
    return nil
}
