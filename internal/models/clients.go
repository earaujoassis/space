package models

import (
	"encoding/json"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/utils"
)

// Client is the client application model/struct
type Client struct {
	Model
	UUID         string         `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"id"`
	Name         string         `gorm:"not null;unique;index" validate:"required,min=3,max=50" json:"name"`
	Description  string         `json:"description"`
	Key          string         `gorm:"not null;unique;index" json:"-"`
	Secret       string         `gorm:"not null" json:"-"`
	Scopes       string         `gorm:"not null" validate:"required,scope_restrict" json:"scopes"`
	CanonicalURI pq.StringArray `gorm:"type:text[];not null" validate:"required,canonical_uri" json:"uri"`
	RedirectURI  pq.StringArray `gorm:"type:text[];not null" validate:"required,redirect_uri" json:"redirect"`
	Type         string         `gorm:"not null" validate:"required,client" json:"-"`
}

// Authentic checks if a secret is valid for a given Client
func (client *Client) Authentic(secret string) bool {
	validSecret := bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secret)) == nil
	return validSecret
}

// SetSecret updates an Client's secret
func (client *Client) SetSecret(secret string) error {
	crypted, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err == nil {
		client.Secret = string(crypted)
		return nil
	}
	return err
}

// BeforeCreate Client model/struct hook
func (client *Client) BeforeCreate(tx *gorm.DB) error {
	client.UUID = generateUUID()
	client.Key = GenerateRandomString(32)
	client.Secret = GenerateRandomString(64)
	if crypted, err := bcrypt.GenerateFromPassword([]byte(client.Secret), bcrypt.DefaultCost); err == nil {
		client.Secret = string(crypted)
	} else {
		return err
	}
	return nil
}

// BeforeSave Client model/struct hook
func (client *Client) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", client)
}

// DefaultCanonicalURI gets the default (first) canonical URI/URL for a client application
func (client *Client) DefaultCanonicalURI() string {
	return client.CanonicalURI[0]
}

// DefaultRedirectURI gets the default (first) redirect URI/URL for a client application
func (client *Client) DefaultRedirectURI() string {
	return client.RedirectURI[0]
}

func (client Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Scopes       string `json:"scopes"`
		CanonicalURI string `json:"uri"`
		RedirectURI  string `json:"redirect"`
	}{
		Id:           client.UUID,
		Name:         client.Name,
		Description:  client.Description,
		Scopes:       client.Scopes,
		CanonicalURI: strings.Join(client.CanonicalURI, "\n"),
		RedirectURI:  strings.Join(client.RedirectURI, "\n"),
	})
}

func (client *Client) HasScope(scope string) bool {
	return strings.Contains(client.Scopes, scope)
}

func (client *Client) HasRequestedScopes(requestedScopes []string) bool {
	validScopes := utils.Scopes(client.Scopes)
	validSet := make(map[string]bool)
	for _, scope := range validScopes {
		validSet[scope] = true
	}

	for _, requested := range requestedScopes {
		if !validSet[requested] {
			return false
		}
	}

	return true
}
