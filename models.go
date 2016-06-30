package main

import (
    "math/rand"
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
    "github.com/satori/go.uuid"
    "gopkg.in/bluesuncorp/validator.v5"
)

type Model struct {
    ID uint                     `gorm:"primary_key" json:"-"`
    CreatedAt time.Time         `gorm:"not null" json:"-"`
    UpdatedAt time.Time         `json:"-"`
}

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
    scope.SetColumn("Key", randStringBytesMaskImprSrc(32))
    scope.SetColumn("Secret", randStringBytesMaskImprSrc(64))
    return nil
}

type Language struct {
    Model
    Name string                 `gorm:"not null;unique;index" validate:"required,min=3"`
    IsoCode string              `gorm:"not null;unique" validate:"required,min=2,max=5"`
}

type User struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"-"`
    PublicId string             `gorm:"not null;unique;index" json:"public_id"`
    Username string             `gorm:"not null;unique;index" validate:"required,alphanum,max=60" json:"username"`
    FirstName string            `gorm:"not null" validate:"required,min=3,max=20" json:"first_name"`
    LastName string             `gorm:"not null" validate:"required,min=3,max=20" json:"last_name"`
    Email string                `gorm:"not null;unique;index" validate:"required,email" json:"email"`
    Passphrase string           `gorm:"not null" validate:"required,min=10" json:"-"`
    Active bool                 `gorm:"default:false" json:"active"`
    Admin bool                  `gorm:"default:false" json:"admin"`
    Client Client               `gorm:"not null" validate:"exists" json:"-"`
    ClientID uint               `gorm:"not null" json:"-"`
    Language Language           `gorm:"not null" validate:"exists" json:"-"`
    LanguageID uint             `gorm:"not null" json:"-"`
    TimezoneIdentifier string   `gorm:"not null;default:'GMT'" json:"timezone_identifier"`
}

type Session struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4"`
    Owner User                  `gorm:"not null" validate:"exists"`
    Maintainer Client           `gorm:"not null" validate:"exists"`
    Moment int                  `gorm:"not null" validate:"required"`
    Ip string                   `gorm:"not null;index" validate:"required"`
    UserAgent string            `gorm:"not null"`
    Invalidated bool            `gorm:"not null"`
}

func (user *User) Authentic(password []byte) bool {
    return bcrypt.CompareHashAndPassword([]byte(user.Passphrase), password) == nil
}

func (user *User) UpdatePassword(password []byte) error {
    pw, err := bcrypt.GenerateFromPassword([]byte(user.Passphrase), bcrypt.DefaultCost)
    if err == nil {
        user.Passphrase = string(pw)
        return nil
    }
    return err
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

const (
    letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImprSrc(n int) string {
    b := make([]byte, n)
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}

func generateUUID() string {
    return uuid.NewV4().String()
}
