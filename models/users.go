package models

import (
    "fmt"
    "strings"

    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
    "github.com/pquerna/otp"
    "github.com/pquerna/otp/totp"

    "github.com/earaujoassis/space/security"
)

// User model/struct
type User struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    PublicID string             `gorm:"not null;unique;index" json:"public_id"`
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

// Authentic checks if a password + passcode combination is valid for a given User
func (user *User) Authentic(password, passcode string) bool {
    var validPasscode bool
    validPassword := bcrypt.CompareHashAndPassword([]byte(user.Passphrase), []byte(password)) == nil
    codeSecret, err := security.Decrypt(defaultKey(), user.CodeSecret)
    if err != nil {
        return false
    }

    validPasscode = totp.Validate(passcode, string(codeSecret))
    return validPasscode && validPassword
}

// UpdatePassword updates an User's password
func (user *User) UpdatePassword(password string) error {
    crypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err == nil {
        user.Passphrase = string(crypted)
        return nil
    }
    return err
}

// GenerateCodeSecret generates a code secret for an user, in order to generate and validate passcodes
func (user *User) GenerateCodeSecret() *otp.Key {
    key, err := totp.Generate(totp.GenerateOpts{
        Issuer:      "quatroLABS.com",
        AccountName: user.Username,
    })
    codeSecret := key.Secret()
    if cryptedCodeSecret, err := security.Encrypt(defaultKey(), []byte(codeSecret)); err == nil {
        user.CodeSecret = string(cryptedCodeSecret)
    } else {
        user.CodeSecret = codeSecret
    }
    if err != nil {
        return nil
    }
    return key
}

// GenerateRecoverSecret generates a recover secret string for an user
func (user *User) GenerateRecoverSecret() (string, error) {
    var secret = strings.ToUpper(fmt.Sprintf("%s-%s-%s-%s",
        GenerateRandomString(4),
        GenerateRandomString(4),
        GenerateRandomString(4),
        GenerateRandomString(4),))
    if cryptedRecoverSecret, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost); err == nil {
        user.RecoverSecret = string(cryptedRecoverSecret)
    } else {
        return secret, err
    }
    return secret, nil
}

// BeforeSave User model/struct hook
func (user *User) BeforeSave(scope *gorm.Scope) error {
    return validateModel("validate", user)
}

// BeforeCreate User model/struct hook
func (user *User) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("PublicID", GenerateRandomString(32))
    if cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Passphrase), bcrypt.DefaultCost); err == nil {
        scope.SetColumn("Passphrase", cryptedPassword)
    } else {
        return err
    }
    return nil
}
