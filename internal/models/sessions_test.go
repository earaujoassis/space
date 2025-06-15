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
	assert.False(t, IsValid("validate", session), "should return false for invalid session")
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
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 10),
	}
	user.Client = client
	user.Language = Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	val := user.GenerateCodeSecret()
	assert.NotNil(t, val)
	_, err = user.GenerateRecoverSecret()
	assert.Nil(t, err, fmt.Sprintf("%s", err))
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
	assert.True(t, IsValid("validate", session), "should return true for valid session")
	err = validateModel("validate", session)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}

func TestHasValidScopes(t *testing.T) {
	assert.False(t, HasValidScopes([]string{WriteScope}))
	assert.True(t, HasValidScopes([]string{PublicScope, OpenIDScope}))
	assert.True(t, HasValidScopes([]string{PublicScope, OpenIDScope, ProfileScope}))
	assert.True(t, HasValidScopes([]string{PublicScope, ReadScope, OpenIDScope, ProfileScope}))
	assert.False(t, HasValidScopes([]string{PublicScope, ReadScope, WriteScope, OpenIDScope, ProfileScope}))
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
