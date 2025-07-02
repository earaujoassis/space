package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/earaujoassis/space/internal/utils"
)

func TestSetSecret(t *testing.T) {
	secret := GenerateRandomString(64)

	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.True(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))

	assert.False(t, client.Authentic(secret))

	client.SetSecret(secret)
	assert.True(t, client.Authentic(secret))
}

func TestDefaultValues(t *testing.T) {
	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.True(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))

	assert.Equal(t, "https://localhost:5000", client.DefaultCanonicalURI())
	assert.Equal(t, "https://localhost:5000/callback", client.DefaultRedirectURI())
}

func TestValidClientModel(t *testing.T) {
	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.True(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
}

func TestValidClientModelWithMultipleScopes(t *testing.T) {
	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       "openid profile public read",
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         ConfidentialClient,
	}

	assert.True(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
}

func TestInvalidClientMissingRequiredFields(t *testing.T) {
	client := Client{
		CanonicalURI: []string{""},
		RedirectURI:  []string{""},
	}

	assert.False(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.NotNil(t, err)
}

func TestInvalidURIClientModel(t *testing.T) {
	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:4000/callback"},
		Type:         PublicClient,
	}

	assert.False(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.NotNil(t, err)
	err = client.BeforeSave(nil)
	assert.NotNil(t, err)
}

func TestInvalidCanonicalURIClientModel(t *testing.T) {
	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       PublicScope,
		CanonicalURI: []string{"ftp://localhost:5000"},
		RedirectURI:  []string{"ftp://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.False(t, IsValid("validate", client))
	val := validateModel("validate", client)
	assert.NotNil(t, val)
	err := client.BeforeSave(nil)
	assert.NotNil(t, err)
}

func TestHasRequestedScopes(t *testing.T) {
	client := Client{
		Scopes: PublicScope,
	}

	assert.True(t, client.HasRequestedScopes([]string{PublicScope}))
	assert.False(t, client.HasRequestedScopes([]string{OpenIDScope}))

	client = Client{
		Scopes: strings.Join([]string{ReadScope, OpenIDScope}, " "),
	}

	assert.False(t, client.HasRequestedScopes([]string{PublicScope}))
	assert.False(t, client.HasRequestedScopes([]string{WriteScope}))
	assert.True(t, client.HasRequestedScopes([]string{OpenIDScope, ReadScope}))
	assert.True(t, client.HasRequestedScopes([]string{ReadScope, OpenIDScope}))
	scope := strings.Join([]string{ReadScope, OpenIDScope}, " ")
	assert.True(t, client.HasRequestedScopes(utils.Scopes(scope)))
}

func TestPublicClientRestrictions(t *testing.T) {
	client := Client{
		Type: PublicClient,
	}

	err := validateModel("validate", client)
	message := fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'required' tag")

	client.Scopes = "public"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'required' tag")

	client.Scopes = "public read"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")
	client.Scopes = "public write"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")
	client.Scopes = "public profile"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")
	client.Scopes = "public email"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")

	client.Scopes = "openid"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")

	client.Scopes = "openid profile"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")
	client.Scopes = "openid email"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")

	client.Scopes = "public openid"
	err = validateModel("validate", client)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Client.Scopes' Error:Field validation for 'Scopes' failed on the 'restrict' tag")
}

func TestClientHasScope(t *testing.T) {
	client := Client{
		Scopes: "public read",
	}
	assert.True(t, client.HasScope("public"))
	assert.True(t, client.HasScope("read"))
	assert.False(t, client.HasScope("openid"))
	assert.False(t, client.HasScope("write"))
}

func TestClientMarshalJSON(t *testing.T) {
	client := Client{
		Name:         "internal",
		Description:  "test internal model only",
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.True(t, IsValid("validate", client))
	err := validateModel("validate", client)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = client.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))

	jsonData, err := json.Marshal(client)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "\"name\":\"internal\"")
}
