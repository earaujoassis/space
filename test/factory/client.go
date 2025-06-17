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

func (f *TestRepositoryFactory) NewClient() *Client {
	client := models.Client{
		Name:         gofakeit.Company(),
		Scopes:       models.PublicScope,
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Type:         models.ConfidentialClient,
	}
	f.manager.Clients().Create(&client)
	clientSecret := models.GenerateRandomString(64)
	client.SetSecret(clientSecret)
	f.manager.Clients().Save(&client)
	localClient := Client{
		Name:   client.Name,
		Key:    client.Key,
		Secret: clientSecret,
		Model:  client,
	}
	return &localClient
}

func (f *TestRepositoryFactory) NewClientWithScopes(scopes string) *Client {
	client := models.Client{
		Name:         gofakeit.Company(),
		Scopes:       scopes,
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Type:         models.ConfidentialClient,
	}
	f.manager.Clients().Create(&client)
	clientSecret := models.GenerateRandomString(64)
	client.SetSecret(clientSecret)
	f.manager.Clients().Save(&client)
	localClient := Client{
		Name:   client.Name,
		Key:    client.Key,
		Secret: clientSecret,
		Model:  client,
	}
	return &localClient
}

func (c *Client) BasicAuthEncode() string {
	return shared.BasicAuthEncode(c.Key, c.Secret)
}
