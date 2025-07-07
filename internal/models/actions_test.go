package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestValidActionModel(t *testing.T) {
	var err error
	var action Action

	action = Action{}
	assert.False(t, IsValid("validate", action))
	err = validateModel("validate", action)
	assert.NotNil(t, err)

	client := Client{
		Model:        Model{ID: 1},
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
		Model:         Model{ID: 1},
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
	action = Action{
		User:        user,
		Client:      client,
		UUID:        generateUUID(),
		Token:       GenerateRandomString(64),
		Moment:      time.Now().UTC().Unix(),
		ExpiresIn:   shortestExpirationLength,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      ReadScope,
		Description: NotSpecialAction,
	}
	action.BeforeSave()
	assert.True(t, IsValid("validate", action))
	err = validateModel("validate", action)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	err = action.Validate()
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}

func TestActionWithinExpirationWindow(t *testing.T) {
	var action Action = Action{
		ExpiresIn: shortestExpirationLength,
	}

	timeTravel := time.Duration(shortestExpirationLength)
	action.Moment = time.Now().UTC().Add(-(timeTravel * time.Second)).Unix()
	assert.True(t, action.WithinExpirationWindow())
	timeTravel = time.Duration(shortestExpirationLength + 1)
	action.Moment = time.Now().UTC().Add(-(timeTravel * time.Second)).Unix()
	assert.False(t, action.WithinExpirationWindow())
}

func TestActionCanUpdateUser(t *testing.T) {
	action := Action{}
	assert.False(t, action.CanUpdateUser())

	action = Action{Description: UpdateUserAction}
	assert.True(t, action.CanUpdateUser())

	action = Action{Description: NotSpecialAction}
	assert.False(t, action.CanUpdateUser())
}

func TestActionGrantsReadAbility(t *testing.T) {
	action := Action{}
	assert.False(t, action.GrantsReadAbility())

	action = Action{Scopes: ReadScope}
	assert.True(t, action.GrantsReadAbility())

	action = Action{Scopes: OpenIDScope}
	assert.True(t, action.GrantsReadAbility())

	action = Action{Scopes: WriteScope}
	assert.True(t, action.GrantsReadAbility())
}

func TestActionGrantsWriteAbility(t *testing.T) {
	action := Action{}
	assert.False(t, action.GrantsWriteAbility())

	action = Action{Scopes: ReadScope}
	assert.False(t, action.GrantsWriteAbility())

	action = Action{Scopes: WriteScope}
	assert.True(t, action.GrantsWriteAbility())
}
