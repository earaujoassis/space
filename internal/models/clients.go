package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	// PublicClient client type
	PublicClient string = "public"
	// ConfidentialClient client type
	ConfidentialClient string = "confidential"
)

// Client is the client application model/struct
type Client struct {
	Model
	UUID         string         `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"id"`
	Name         string         `gorm:"not null;unique;index" validate:"required,min=3,max=20" json:"name"`
	Description  string         `json:"description"`
	Key          string         `gorm:"not null;unique;index" json:"-"`
	Secret       string         `gorm:"not null" validate:"required" json:"-"`
	Scopes       string         `gorm:"not null" validate:"required,restrict" json:"scopes"`
	CanonicalURI pq.StringArray `gorm:"type:text[];not null" validate:"required,canonical" json:"uri"`
	RedirectURI  pq.StringArray `gorm:"type:text[];not null" validate:"required,redirect" json:"redirect"`
	Type         string         `gorm:"not null" validate:"required,client" json:"-"`
}

func validClientScopes(fl validator.FieldLevel) bool {
	scopesField := fl.Field().String()
	// TODO A PublicClient can't have a ReadScope
	if scopesField != PublicScope && scopesField != ReadScope {
		return false
	}
	return true
}

func validClientType(fl validator.FieldLevel) bool {
	typeField := fl.Field().String()
	if typeField != PublicClient && typeField != ConfidentialClient {
		return false
	}
	return true
}

func validCanonicalURIs(fl validator.FieldLevel) bool {
	// It's not a Client model
	if !fl.Top().CanConvert(reflect.TypeOf(Client{})) {
		return true
	}

	currentClient := fl.Top().Interface().(Client)
	// Unfortunately, the Jupiter (internal client) is created with the following value
	if len(currentClient.CanonicalURI) == 1 && currentClient.CanonicalURI[0] == "localhost" {
		return true
	}

	for i := range currentClient.CanonicalURI {
		currentEntry := currentClient.CanonicalURI[i]
		u, err := url.Parse(currentEntry)
		if err != nil {
			return false
		}

		if !strings.Contains(u.Scheme, "http") || u.Path != "" || u.Host == "" {
			return false
		}
	}

	return true
}

func validRedirectURIs(fl validator.FieldLevel) bool {
	// It's not a Client model
	if !fl.Top().CanConvert(reflect.TypeOf(Client{})) {
		return true
	}

	currentClient := fl.Top().Interface().(Client)
	canonicalUri := currentClient.CanonicalURI
	redirectUri := currentClient.RedirectURI

	// Unfortunately, the Jupiter (internal client) is created with the following values
	if len(canonicalUri) == 1 && canonicalUri[0] == "localhost" && len(redirectUri) == 1 && redirectUri[0] == "/" {
		return true
	}

	for i := range redirectUri {
		currentEntry := redirectUri[i]
		u, err := url.Parse(currentEntry)
		if err != nil {
			return false
		}
		currentCanonical := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
		if !slices.Contains(canonicalUri, currentCanonical) {
			return false
		}
	}

	return true
}

// Authentic checks if a secret is valid for a given Client
func (client *Client) Authentic(secret string) bool {
	validSecret := bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secret)) == nil
	return validSecret
}

// UpdateSecret updates an Client's secret
func (client *Client) UpdateSecret(secret string) error {
	crypted, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err == nil {
		client.Secret = string(crypted)
		return nil
	}
	return err
}

// BeforeSave Client model/struct hook
func (client *Client) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", client)
}

// BeforeCreate Client model/struct hook
func (client *Client) BeforeCreate(tx *gorm.DB) error {
	client.UUID = generateUUID()
	client.Key = GenerateRandomString(32)
	if crypted, err := bcrypt.GenerateFromPassword([]byte(client.Secret), bcrypt.DefaultCost); err == nil {
		client.Secret = string(crypted)
	} else {
		return err
	}
	return nil
}

// DefaultRedirectURI gets the default (first) redirect URI/URL for a client application
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
