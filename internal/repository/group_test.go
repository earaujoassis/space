package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestGroupRepository__FindOrCreate() {
	repository := NewGroupRepository(s.db)
	clients := NewClientRepository(s.db)
	languages := NewLanguageRepository(s.db)
	users := NewUserRepository(s.db)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	err := clients.Create(&client)
	s.Require().NoError(err)
	s.Require().NotZero(client.ID)
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	err = languages.Create(&language)
	s.Require().NoError(err)
	s.Require().NotZero(language.ID)
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	err = users.Create(&user)
	s.Require().NoError(err)
	s.Require().NotZero(user.ID)
	group := models.Group{
		User:   user,
		Client: client,
		Tags:   []string{"testing"},
	}
	err = repository.Create(&group)
	s.Require().NoError(err)
	s.Require().NotZero(group.ID)

	retrieved := repository.FindOrCreate(user, client)
	s.Require().NotZero(retrieved.ID)
	s.Equal(client.Name, retrieved.Client.Name)
	s.Equal(user.Username, retrieved.User.Username)
	s.Equal(group.ID, retrieved.ID)
}
