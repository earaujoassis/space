package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Group struct {
	Model
	User     User           `gorm:"not null;foreignKey:UserID" validate:"nostructlevel" json:"-"`
	UserID   uint           `gorm:"not null" json:"-"`
	Client   Client         `gorm:"not null;foreignKey:ClientID" validate:"nostructlevel" json:"-"`
	ClientID uint           `gorm:"not null" json:"-"`
	Tags     pq.StringArray `gorm:"type:text[];not null" validate:"required" json:"groups"`
}

// BeforeSave Language model/struct hook
func (group *Group) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", group)
}
