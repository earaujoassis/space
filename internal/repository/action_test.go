package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestActionRepository__CreationAndRetrieval() {
	repository := NewActionRepository(s.ms)
	clients := NewClientRepository(s.db)
	languages := NewLanguageRepository(s.db)
	users := NewUserRepository(s.db)

	user := models.User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 32),
	}
	users.SetRecoverSecret(&user)
	users.SetCodeSecret(&user).Secret()
	user.Client = clients.FindOrCreate(models.DefaultClient)
	user.Language = languages.FindOrCreate("English", "en-US")
	users.Create(&user)

	actionToken := models.Action{
		User:        user,
		Client:      user.Client,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      models.WriteScope,
		Description: models.NotSpecialAction,
	}
	repository.Create(&actionToken)
	s.Equal(len(actionToken.Token), 64)
	retrievedByToken := repository.FindByToken(actionToken.Token)
	s.Equal(retrievedByToken.Token, actionToken.Token)
	s.Equal(retrievedByToken.UUID, actionToken.UUID)
	retrievedByUuid := repository.FindByUUID(actionToken.UUID)
	s.Equal(retrievedByUuid.Token, actionToken.Token)
	s.Equal(retrievedByUuid.UUID, actionToken.UUID)
}

func (s *RepositoryTestSuite) TestActionRepository__Authentication() {
	repository := NewActionRepository(s.ms)
	clients := NewClientRepository(s.db)
	languages := NewLanguageRepository(s.db)
	users := NewUserRepository(s.db)

	user := models.User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 32),
	}
	users.SetRecoverSecret(&user)
	users.SetCodeSecret(&user).Secret()
	user.Client = clients.FindOrCreate(models.DefaultClient)
	user.Language = languages.FindOrCreate("English", "en-US")
	users.Create(&user)

	actionToken := models.Action{
		User:        user,
		Client:      user.Client,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      models.WriteScope,
		Description: models.NotSpecialAction,
	}
	repository.Create(&actionToken)
	s.Equal(len(actionToken.Token), 64)
	authenticated := repository.Authentication(actionToken.Token)
	s.Equal(len(authenticated.Token), 64)
	s.Equal(authenticated.Token, actionToken.Token)
	s.Equal(authenticated.UUID, actionToken.UUID)
	s.True(authenticated.WithinExpirationWindow())
}

func (s *RepositoryTestSuite) TestActionRepository__Delete() {
	repository := NewActionRepository(s.ms)
	clients := NewClientRepository(s.db)
	languages := NewLanguageRepository(s.db)
	users := NewUserRepository(s.db)

	user := models.User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: gofakeit.Password(true, true, true, true, false, 32),
	}
	users.SetRecoverSecret(&user)
	users.SetCodeSecret(&user).Secret()
	user.Client = clients.FindOrCreate(models.DefaultClient)
	user.Language = languages.FindOrCreate("English", "en-US")
	users.Create(&user)

	actionToken := models.Action{
		User:        user,
		Client:      user.Client,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      models.WriteScope,
		Description: models.NotSpecialAction,
	}
	repository.Create(&actionToken)
	s.Equal(len(actionToken.Token), 64)
	token := actionToken.Token
	repository.Delete(actionToken)
	retrieved := repository.FindByToken(token)
	s.Empty(retrieved.UUID)
	s.Empty(retrieved.Token)
}
