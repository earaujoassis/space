package models

import (
	"gorm.io/gorm"
)

// Language model/struct represents a Language option through the Application UI
type Language struct {
	Model
	Name    string `gorm:"not null;unique;index" validate:"required,min=2"`
	IsoCode string `gorm:"not null;unique" validate:"required,min=2,max=5"`
}

// BeforeSave Language model/struct hook
func (language *Language) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", language)
}
