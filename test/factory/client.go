package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
)

type Client struct {
	Name   string
	Key    string
	Secret string
	Model  models.Client
}

func (c *Client) BasicAuthEncode() string {
	return shared.BasicAuthEncode(c.Key, c.Secret)
}

func (f *TestRepositoryFactory) NewClient() *Client {
	repositories := f.manager
	client := models.Client{
		Name:         gofakeit.Company(),
		Scopes:       models.PublicScope,
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Type:         models.ConfidentialClient,
	}
	repositories.Clients().Create(&client)
	clientSecret := models.GenerateRandomString(64)
	client.SetSecret(clientSecret)
	repositories.Clients().Save(&client)
	localClient := Client{
		Name:   client.Name,
		Key:    client.Key,
		Secret: clientSecret,
		Model:  client,
	}
	return &localClient
}

func (f *TestRepositoryFactory) NewClientWithScopes(scopes string) *Client {
	repositories := f.manager
	client := models.Client{
		Name:         gofakeit.Company(),
		Scopes:       scopes,
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Type:         models.ConfidentialClient,
	}
	repositories.Clients().Create(&client)
	clientSecret := models.GenerateRandomString(64)
	client.SetSecret(clientSecret)
	repositories.Clients().Save(&client)
	localClient := Client{
		Name:   client.Name,
		Key:    client.Key,
		Secret: clientSecret,
		Model:  client,
	}
	return &localClient
}

func (f *TestRepositoryFactory) DefaultClient() *Client {
	client := f.manager.Clients().FindOrCreate(models.DefaultClient)
	localClient := Client{
		Name:  client.Name,
		Key:   client.Key,
		Model: client,
	}
	return &localClient
}
