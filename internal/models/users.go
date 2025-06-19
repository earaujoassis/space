package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model/struct
type User struct {
	Model
	UUID               string   `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
	PublicID           string   `gorm:"not null;unique;index" json:"public_id"`
	Username           string   `gorm:"not null;unique;index" validate:"required,alphanum,max=60" json:"-"`
	FirstName          string   `gorm:"not null" validate:"required,min=3,max=20" essential:"required,min=3,max=20" json:"first_name"`
	LastName           string   `gorm:"not null" validate:"required,min=3,max=20" essential:"required,min=3,max=20" json:"last_name"`
	Email              string   `gorm:"not null;unique;index" validate:"required,email" essential:"required,email" json:"email"`
	EmailVerified      bool     `gorm:"not null;default:false" json:"email_verified"`
	Passphrase         string   `gorm:"not null" validate:"required" essential:"required,min=10" json:"-"`
	Active             bool     `gorm:"not null;default:false" json:"active"`
	Admin              bool     `gorm:"not null;default:false" json:"-"`
	Client             Client   `gorm:"not null" validate:"required" json:"-"`
	ClientID           uint     `gorm:"not null" json:"-"`
	Language           Language `gorm:"not null" validate:"required" json:"-"`
	LanguageID         uint     `gorm:"not null" json:"-"`
	TimezoneIdentifier string   `gorm:"not null;default:'GMT'" json:"timezone_identifier"`
	CodeSecret         string   `gorm:"not null" validate:"required" json:"-"`
	RecoverSecret      string   `gorm:"not null" validate:"required" json:"-"`
}

// BeforeSave User model/struct hook
func (user *User) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", user)
}

// BeforeCreate User model/struct hook
func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.UUID = generateUUID()
	user.PublicID = GenerateRandomString(32)
	if cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Passphrase), bcrypt.DefaultCost); err == nil {
		user.Passphrase = string(cryptedPassword)
	} else {
		return err
	}
	return nil
}
