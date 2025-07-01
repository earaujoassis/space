package models

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestValidEmailModel(t *testing.T) {
	email := Email{
		Address: "",
	}

	assert.False(t, IsValid("validate", email))
	val := validateModel("validate", email)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))
	err := email.BeforeSave(nil)
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
	assert.Nil(t, err, fmt.Sprintf("%s", err))
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
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	email = Email{
		User:    user,
		Address: gofakeit.Email(),
	}

	err = email.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	assert.True(t, IsValid("validate", email))
	val = validateModel("validate", email)
	assert.Nil(t, val, fmt.Sprintf("%v", val))
	err = email.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
}
