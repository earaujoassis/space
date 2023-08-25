package models

import (
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"

	"github.com/earaujoassis/space/internal/config"
)

// Model is the base model/struct for any model in the application/system
type Model struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `json:"-"`
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateRandomString returns a random string with `n` as the length
func GenerateRandomString(n int) string {
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

func validateModel(tagName string, model interface{}) error {
	validate := validator.New()
	validate.SetTagName(tagName)
	validate.RegisterValidation("client", validClientType)
	validate.RegisterValidation("scope", validScope)
	validate.RegisterValidation("restrict", validClientScopes)
	validate.RegisterValidation("token", validTokenType)
	validate.RegisterValidation("canonical", validCanonicalURIs)
	validate.RegisterValidation("redirect", validRedirectURIs)
	validate.RegisterValidation("action", validAction)
	err := validate.Struct(model)
	if err != nil {
		return err
	}
	return nil
}

// IsValid checks if a `model` entry is valid, given the `tagName` (scope) for validation
func IsValid(tagName string, model interface{}) bool {
	err := validateModel(tagName, model)
	if err != nil {
		return false
	}
	return true
}

func defaultKey() []byte {
	keyString := config.GetGlobalConfig().StorageSecret
	return []byte(keyString)
}

func (model Model) IsNewRecord() bool {
	return model.ID == 0
}
