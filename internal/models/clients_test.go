package models

import (
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

	assert.True(t, IsValid("validate", client), "should return true for the valid client")
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
