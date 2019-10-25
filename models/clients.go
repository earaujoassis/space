package models

import (
    "fmt"
    "strings"
    "net/url"

    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
)

const (
    // PublicClient client type
    PublicClient        string = "public"
    // ConfidentialClient client type
    ConfidentialClient  string = "confidential"
)

// Client is the client application model/struct
type Client struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"id"`
    Name string                 `gorm:"not null;unique;index" validate:"required,min=3,max=20" json:"name"`
    Description string          `json:"description"`
    Key string                  `gorm:"not null;unique;index" json:"-"`
    Secret string               `gorm:"not null" validate:"required" json:"-"`
    Scopes string               `gorm:"not null" validate:"required" json:"-"`
    CanonicalURI string         `gorm:"not null" validate:"required,canonical" json:"uri"`
    RedirectURI string          `gorm:"not null" validate:"required,redirect" json:"redirect"`
    Type string                 `gorm:"not null" validate:"required,client" json:"-"`
}

func validClientType(top interface{}, current interface{}, field interface{}, param string) bool {
    typeField := field.(string)
    if typeField != PublicClient && typeField != ConfidentialClient {
        return false
    }
    return true
}

func validCanonicalURIs(top interface{}, current interface{}, field interface{}, param string) bool {
    canonicalURIField := field.(string)

    // Unfortunately, the Jupiter (internal client) is created with the following value
    if canonicalURIField == "localhost" {
        return true
    }

    entries := strings.Split(canonicalURIField, "\n")
    for i := range entries {
        currentEntry := entries[i]
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

func validRedirectURIs(top interface{}, current interface{}, field interface{}, param string) bool {
    currentClient, ok := current.(Client)

    // It's not a Client model
    if !ok {
        return true
    }

    canonicalURI := currentClient.CanonicalURI
    redirectURI := currentClient.RedirectURI

    // Unfortunately, the Jupiter (internal client) is created with the following values
    if canonicalURI == "localhost" && redirectURI == "/" {
        return true
    }

    entries := strings.Split(redirectURI, "\n")
    for i := range entries {
        currentEntry := entries[i]
        u, err := url.Parse(currentEntry)
        if err != nil {
            return false
        }
        currentCanonical := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
        if !strings.Contains(canonicalURI, currentCanonical) {
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
func (client *Client) BeforeSave(scope *gorm.Scope) error {
    return validateModel("validate", client)
}

// BeforeCreate Client model/struct hook
func (client *Client) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("UUID", generateUUID())
    scope.SetColumn("Key", GenerateRandomString(32))
    if crypted, err := bcrypt.GenerateFromPassword([]byte(client.Secret), bcrypt.DefaultCost); err == nil {
        scope.SetColumn("Secret", crypted)
    } else {
        return err
    }
    return nil
}

// DefaultRedirectURI gets the default (first) redirect URI/URL for a client application
func (client *Client) DefaultRedirectURI() string {
    return strings.Split(client.RedirectURI, "\n")[0]
}
