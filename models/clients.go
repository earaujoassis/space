package models

import (
    "github.com/jinzhu/gorm"
    "gopkg.in/bluesuncorp/validator.v5"
)

type Client struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    Name string                 `gorm:"not null;unique;index" validate:"required,min=3,max=20" json:"name"`
    Description string          `json:"description"`
    Key string                  `gorm:"not null;unique;index" validate:"required" json:"-"`
    Secret string               `gorm:"not null" validate:"required" json:"-"`
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
    //scope.SetColumn("PublicId", randStringBytesMaskImprSrc(32))
    scope.SetColumn("Key", randStringBytesMaskImprSrc(32))
    scope.SetColumn("Secret", randStringBytesMaskImprSrc(64))
    return nil
}
