package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestClientRepository__FindOrCreate() {
	repository := NewClientRepository(s.DB)

	defaultClient := repository.FindOrCreate(models.DefaultClient)
	s.Require().NotZero(defaultClient.ID)
	s.Equal(defaultClient.Name, "Jupiter")
	s.Equal(defaultClient.Scopes, models.PublicScope)
	s.Equal(defaultClient.Type, models.PublicClient)

	clientName := gofakeit.Company()
	client := models.Client{
		Name:         clientName,
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	err := repository.Create(&client)
	s.Require().NoError(err)
	retrieved := repository.FindOrCreate(clientName)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.Name, clientName)
	s.Equal(retrieved.Type, models.ConfidentialClient)
}

func (s *RepositoryTestSuite) TestClientRepository__FindByKey() {
	repository := NewClientRepository(s.DB)

	clientName := gofakeit.Company()
	client := models.Client{
		Name:         clientName,
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	err := repository.Create(&client)
	s.Require().NoError(err)
	key := client.Key
	retrieved := repository.FindByKey(key)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.Name, clientName)
	s.Equal(retrieved.Type, models.ConfidentialClient)
	s.Equal(retrieved.Key, key)
}

func (s *RepositoryTestSuite) TestClientRepository__FindByUUID() {
	repository := NewClientRepository(s.DB)

	clientName := gofakeit.Company()
	client := models.Client{
		Name:         clientName,
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	err := repository.Create(&client)
	s.Require().NoError(err)
	uuid := client.UUID
	retrieved := repository.FindByUUID(uuid)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.Name, clientName)
	s.Equal(retrieved.Type, models.ConfidentialClient)
	s.Equal(retrieved.UUID, uuid)
}

func (s *RepositoryTestSuite) TestClientRepository__Authentication() {
	repository := NewClientRepository(s.DB)

	clientName := gofakeit.Company()
	client := models.Client{
		Name:         clientName,
		Scopes:       models.PublicScope,
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Type:         models.ConfidentialClient,
	}
	repository.Create(&client)
	clientSecret := models.GenerateRandomString(64)
	clientKey := client.Key
	client.SetSecret(clientSecret)
	repository.Save(&client)

	retrieved := repository.Authentication(clientKey, clientSecret)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.Name, clientName)
	s.Equal(retrieved.Key, clientKey)
	s.Equal(retrieved.Type, models.ConfidentialClient)

	anotherSecret := models.GenerateRandomString(64)
	anotherClient := repository.Authentication(clientKey, anotherSecret)
	s.Zero(anotherClient.ID)
}

func (s *RepositoryTestSuite) TestClientRepository__GetActive() {
	repository := NewClientRepository(s.DB)

	defaultClient := repository.FindOrCreate(models.DefaultClient)
	s.Require().NotZero(defaultClient.ID)
	s.Equal(defaultClient.Name, "Jupiter")
	s.Equal(defaultClient.Scopes, models.PublicScope)
	s.Equal(defaultClient.Type, models.PublicClient)

	clientName := gofakeit.Company()
	client := models.Client{
		Name:         clientName,
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	err := repository.Create(&client)
	s.Require().NoError(err)

	activeClients := repository.GetActive()
	s.Require().Equal(len(activeClients), 1)
	retrieved := activeClients[0]
	s.Equal(retrieved.Name, clientName)
}
