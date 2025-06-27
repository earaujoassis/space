package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestValidSessionModel(t *testing.T) {
	var err error
	var session Session

	session = Session{}
	assert.False(t, IsValid("validate", session))
	err = validateModel("validate", session)
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
	session = Session{
		User:      user,
		Client:    client,
		UUID:      generateUUID(),
		Token:     GenerateRandomString(64),
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    ReadScope,
		TokenType: RefreshToken,
	}
	err = session.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	assert.True(t, IsValid("validate", session))
	err = validateModel("validate", session)
	assert.Nil(t, err, fmt.Sprintf("%s", err))

	err = session.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	assert.True(t, IsValid("validate", session))
	err = validateModel("validate", session)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}

func TestSessionWithinExpirationWindow(t *testing.T) {
	var session Session = Session{
		ExpiresIn: refreshableExpirationLength,
	}

	timeTravel := time.Duration(refreshableExpirationLength)
	session.Moment = time.Now().UTC().Add(-(timeTravel * time.Second)).Unix()
	assert.True(t, session.WithinExpirationWindow())
	timeTravel = time.Duration(refreshableExpirationLength + 1)
	session.Moment = time.Now().UTC().Add(-(timeTravel * time.Second)).Unix()
	assert.False(t, session.WithinExpirationWindow())
}

func TestValidTokenType(t *testing.T) {
	var err error
	var message string

	session := Session{}

	session.TokenType = "application_token"
	err = validateModel("validate", session)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Session.TokenType' Error:Field validation for 'TokenType' failed on the 'token' tag")

	session.TokenType = "access_token"
	err = validateModel("validate", session)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Session.TokenType' Error:Field validation for 'TokenType' failed on the 'token' tag")

	session.TokenType = "refresh_token"
	err = validateModel("validate", session)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Session.TokenType' Error:Field validation for 'TokenType' failed on the 'token' tag")

	session.TokenType = "grant_token"
	err = validateModel("validate", session)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Session.TokenType' Error:Field validation for 'TokenType' failed on the 'token' tag")

	session.TokenType = "id_token"
	err = validateModel("validate", session)
	message = fmt.Sprintf("%s", err)
	assert.NotContains(t, message, "Key: 'Session.TokenType' Error:Field validation for 'TokenType' failed on the 'token' tag")

	session.TokenType = "invalid_token"
	err = validateModel("validate", session)
	message = fmt.Sprintf("%s", err)
	assert.Contains(t, message, "Key: 'Session.TokenType' Error:Field validation for 'TokenType' failed on the 'token' tag")
}

func TestSessionGrantsReadAbility(t *testing.T) {
	session := Session{}
	assert.False(t, session.GrantsReadAbility())

	session = Session{Scopes: ReadScope}
	assert.True(t, session.GrantsReadAbility())

	session = Session{Scopes: OpenIDScope}
	assert.True(t, session.GrantsReadAbility())

	session = Session{Scopes: WriteScope}
	assert.True(t, session.GrantsReadAbility())
}

func TestSessionGrantsWriteAbility(t *testing.T) {
	session := Session{}
	assert.False(t, session.GrantsWriteAbility())

	session = Session{Scopes: ReadScope}
	assert.False(t, session.GrantsWriteAbility())

	session = Session{Scopes: WriteScope}
	assert.True(t, session.GrantsWriteAbility())
}
