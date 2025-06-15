package models

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidClientModel(t *testing.T) {
	var client Client = Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.True(t, IsValid("validate", client), "should return true for valid client")
}

func TestValidClientModelWithMultipleScopes(t *testing.T) {
	var client Client = Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       "openid profile public read",
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.True(t, IsValid("validate", client), "should return true for valid client")
	err := validateModel("validate", client)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}

func TestInvalidClientMissingRequiredFields(t *testing.T) {
	var client Client = Client{
		Name:         "",
		Description:  "", // this can be empty
		Secret:       "",
		Scopes:       "",
		CanonicalURI: []string{""},
		RedirectURI:  []string{""},
		Type:         "",
	}

	assert.False(t, IsValid("validate", client), "should return false for missing required fields")
}

func TestInvalidURIClientModel(t *testing.T) {
	var client Client = Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       PublicScope,
		CanonicalURI: []string{"https://localhost:5000"},
		RedirectURI:  []string{"https://localhost:4000/callback"},
		Type:         PublicClient,
	}

	assert.False(t, IsValid("validate", client), "should return false for mismatch between canonical and redirect URIs")
}

func TestInvalidCanonicalURIClientModel(t *testing.T) {
	var client Client = Client{
		Name:         "internal",
		Description:  "test internal model only",
		Secret:       GenerateRandomString(64),
		Scopes:       PublicScope,
		CanonicalURI: []string{"ftp://localhost:5000"},
		RedirectURI:  []string{"ftp://localhost:5000/callback"},
		Type:         PublicClient,
	}

	assert.False(t, IsValid("validate", client), "should return false for invalid canonical Scheme/URI")
}

func TestHasRequestedScopes(t *testing.T) {
	var client Client = Client{
		Scopes: PublicScope,
	}

	assert.True(t, client.HasRequestedScopes([]string{ PublicScope }), "should have the requested scope")
	assert.False(t, client.HasRequestedScopes([]string{ OpenIDScope }), "should not have the requested scope")

	client = Client{
		Scopes: strings.Join([]string{ ReadScope, OpenIDScope }, " "),
	}

	assert.False(t, client.HasRequestedScopes([]string{ PublicScope }), "should not have the requested scope")
	assert.False(t, client.HasRequestedScopes([]string{ WriteScope }), "should not have the requested scope")
	assert.True(t, client.HasRequestedScopes([]string{ OpenIDScope, ReadScope }), "should have the requested scope")
	assert.True(t, client.HasRequestedScopes([]string{ ReadScope, OpenIDScope }), "should have the requested scope")
	scope := strings.Join([]string{ ReadScope, OpenIDScope }, " ")
	assert.True(t, client.HasRequestedScopes(strings.Split(scope, " ")), "should have the requested scope")
}
