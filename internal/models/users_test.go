package models

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestValidUserModel(t *testing.T) {
	user := User{}

	assert.False(t, IsValid("validate", user), "should return false for invalid user")
	assert.False(t, IsValid("essential", user), "should return false for invalid user")

	user = User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 32),
	}

	assert.True(t, IsValid("essential", user), "should return true for essential user validation")
	assert.False(t, IsValid("validate", user), "should return false for invalid user")
}

func TestValidUserPassword(t *testing.T) {
	user := User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 9),
	}

	assert.False(t, IsValid("essential", user), "should return false for essential user validation")
	err := validateModel("essential", user)
	message := fmt.Sprintf("%s", err)
	assert.Equal(t, "Key: 'User.Passphrase' Error:Field validation for 'Passphrase' failed on the 'min' tag", message)

	user = User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 10),
	}

	assert.True(t, IsValid("essential", user), "should return true for essential user validation")
	err = validateModel("essential", user)
	assert.Nil(t, err)
}

func TestUserCreation(t *testing.T) {
	var err error

	user := User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 10),
	}

	assert.True(t, IsValid("essential", user), "should return true for essential user validation")
	val := user.GenerateCodeSecret()
	assert.NotNil(t, val)
	_, err = user.GenerateRecoverSecret()
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	err = user.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))

	assert.True(t, IsValid("essential", user), "should return true for essential user validation")
	err = validateModel("essential", user)
	assert.Nil(t, err, fmt.Sprintf("%s", err))

	user.Client = Client{
		Name:         gofakeit.Company(),
		Secret:       GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       PublicScope,
		Type:         PublicClient,
	}
	user.Language = Language{
		Name:    "English",
		IsoCode: "en-US",
	}

	assert.True(t, IsValid("validate", user), "should return true for user validation")
	err = validateModel("validate", user)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	err = user.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}
