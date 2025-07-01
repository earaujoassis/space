package models

import (
	"gorm.io/gorm"
)

// Email model/struct represents an Email option through the Application UI
type Email struct {
	Model
	UUID     string `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"id"`
	User     User   `gorm:"not null;foreignKey:UserID" validate:"required" json:"-"`
	UserID   uint   `gorm:"not null" json:"-"`
	Address  string `gorm:"not null;unique;index" validate:"required,email" json:"address"`
	Verified bool   `gorm:"not null;default:false" json:"verified"`
}

// BeforeSave Email model/struct hook
func (email *Email) BeforeSave(tx *gorm.DB) error {
	email.UUID = generateUUID()
	return validateModel("validate", email)
}
