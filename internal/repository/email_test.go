package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestEmailRepository__Create() {
	repository := NewEmailRepository(s.DB)

	email := models.Email{}
	err := repository.Create(&email)
	s.Error(err)

	client := models.Client{
		Name:         "internal",
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	err = client.BeforeSave(nil)
	s.NoError(err)
	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = client
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err = user.BeforeSave(nil)
	s.NoError(err)
	email = models.Email{
		User:    user,
		Address: gofakeit.Email(),
	}
	err = repository.Create(&email)
	s.NoError(err)
}

func (s *RepositoryTestSuite) TestEmailRepository__GetAllForUser() {
	repository := NewEmailRepository(s.DB)

	client := models.Client{
		Name:         "internal",
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	err := client.BeforeSave(nil)
	s.NoError(err)
	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = client
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err = user.BeforeSave(nil)
	s.NoError(err)
	email := models.Email{
		User:    user,
		Address: gofakeit.Email(),
	}
	err = repository.Create(&email)
	s.Require().NoError(err)
	user = email.User

	email = models.Email{
		User:    user,
		Address: gofakeit.Email(),
	}
	err = repository.Create(&email)
	s.Require().NoError(err)

	emails := repository.GetAllForUser(user)
	s.Equal(len(emails), 2)
}
