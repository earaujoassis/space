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
	action = Action{
		User:        user,
		Client:      client,
		UUID:        generateUUID(),
		CreatedAt:   time.Now().UTC(),
		Token:       GenerateRandomString(64),
		Moment:      time.Now().UTC().Unix(),
		ExpiresIn:   shortestExpirationLength,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      ReadScope,
		Description: NotSpecialAction,
	}
	action.BeforeSave()
	err = validateModel("validate", action)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	assert.True(t, IsValid("validate", action))
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
