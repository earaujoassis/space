package models

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type Setting struct {
	Model
	User     User   `gorm:"not null;foreignKey:UserID" validate:"required" json:"-"`
	UserID   uint   `gorm:"not null" json:"-"`
	Realm    string `gorm:"not null" validate:"required" json:"-"`
	Category string `gorm:"not null" validate:"required" json:"-"`
	Property string `gorm:"not null" validate:"required" json:"-"`
	Type     string `gorm:"not null" validate:"required" json:"type"`
	Value    string `gorm:"type:text" json:"value"`
}

// BeforeSave Language model/struct hook
func (setting *Setting) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", setting)
}

func (setting *Setting) DeserializeValue() (interface{}, error) {
	switch setting.Type {
	case "bool":
		return strconv.ParseBool(setting.Value)
	case "string":
		return setting.Value, nil
	case "int":
		return strconv.ParseInt(setting.Value, 10, 64)
	default:
		return setting.Value, nil
	}
}

func (setting Setting) Reduce() (string, interface{}) {
	key := fmt.Sprintf("%s.%s.%s",
		setting.Realm, setting.Category, setting.Property)
	value, _ := setting.DeserializeValue()
	return key, value
}
