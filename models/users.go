package models

import (
    "fmt"
    "strings"

    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
    "gopkg.in/bluesuncorp/validator.v5"
    "github.com/pquerna/otp"
    "github.com/pquerna/otp/totp"
)

type User struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    PublicId string             `gorm:"not null;unique;index" json:"public_id"`
    Username string             `gorm:"not null;unique;index" validate:"required,alphanum,max=60" json:"-"`
    FirstName string            `gorm:"not null" validate:"required,min=3,max=20" essential:"required,min=3,max=20" json:"first_name"`
    LastName string             `gorm:"not null" validate:"required,min=3,max=20" essential:"required,min=3,max=20" json:"last_name"`
    Email string                `gorm:"not null;unique;index" validate:"required,email" essential:"required,email" json:"email"`
    Passphrase string           `gorm:"not null" validate:"required" essential:"required,min=10" json:"-"`
    Active bool                 `gorm:"not null;default:false" json:"active"`
    Admin bool                  `gorm:"not null;default:false" json:"-"`
    Client Client               `gorm:"not null" validate:"exists" json:"-"`
    ClientID uint               `gorm:"not null" json:"-"`
    Language Language           `gorm:"not null" validate:"exists" json:"-"`
    LanguageID uint             `gorm:"not null" json:"-"`
    TimezoneIdentifier string   `gorm:"not null;default:'GMT'" json:"timezone_identifier"`
    CodeSecret string           `gorm:"not null" validate:"required" json:"-"`
    RecoverSecret string        `gorm:"not null" validate:"required" json:"-"`
}

func (user *User) Authentic(password, passcode string) bool {
    validPassword := bcrypt.CompareHashAndPassword([]byte(user.Passphrase), []byte(password)) == nil
    validPasscode := totp.Validate(passcode, user.CodeSecret)
    return validPasscode && validPassword
}

func (user *User) UpdatePassword(password string) error {
    pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err == nil {
        user.Passphrase = string(pw)
        return nil
    }
    return err
}

func (user *User) GenerateCodeSecret() *otp.Key {
    key, err := totp.Generate(totp.GenerateOpts{
        Issuer:      "QuatroLabs.com",
        AccountName: user.Username,
    })
    user.CodeSecret = key.Secret()
    if err != nil {
        return nil
    }
    return key
}

func (user *User) GenerateRecoverSecret() string {
    var secret string = strings.ToUpper(fmt.Sprintf("%s-%s-%s-%s",
        randStringBytesMaskImprSrc(4),
        randStringBytesMaskImprSrc(4),
        randStringBytesMaskImprSrc(4),
        randStringBytesMaskImprSrc(4),))
    user.RecoverSecret = secret
    return secret
}

func (user *User) BeforeSave(scope *gorm.Scope) error {
    validate := validator.New("validate", validator.BakedInValidators)
    err := validate.Struct(user)
    if err != nil {
        return err
    }
    return nil
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("PublicId", randStringBytesMaskImprSrc(32))
    if pw, err := bcrypt.GenerateFromPassword([]byte(user.Passphrase), bcrypt.DefaultCost); err == nil {
        scope.SetColumn("Passphrase", pw)
    } else {
        return err
    }
    return nil
}
