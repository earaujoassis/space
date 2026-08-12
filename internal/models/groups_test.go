package models

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidGroupModel(t *testing.T) {
	group := Group{}
	assert.False(t, IsValid("validate", group))
	val := validateModel("validate", group)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))
	err := group.BeforeSave(nil)
	assert.NotNil(t, err, fmt.Sprintf("%v", err))

	client := Client{
		Name:         "internal",
		Secret:       GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       PublicScope,
		Type:         PublicClient,
	}
	err = client.BeforeSave(nil)
	require.Nil(t, err, fmt.Sprintf("%s", err))
	user := User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = client
	user.Language = Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err = user.BeforeSave(nil)
	require.Nil(t, err, fmt.Sprintf("%s", err))
	group = Group{
		User:   user,
		Client: client,
		Tags:   []string{"testing"},
	}

	assert.True(t, IsValid("validate", group))
	val = validateModel("validate", group)
	assert.Nil(t, val, fmt.Sprintf("%v", val))
	err = group.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
}
