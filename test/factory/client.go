package factory

import (
	"log"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/shared"
)

type Client struct {
	Name   string
	Key    string
	Secret string
	Model  models.Client
}

func NewClient() *Client {
	client := models.Client{
		Name:         gofakeit.Company(),
		Scopes:       models.PublicScope,
		CanonicalURI: []string{ "http://localhost" },
		RedirectURI:  []string{ "http://localhost/callback" },
		Type:         models.ConfidentialClient,
	}
	ok, err := services.CreateNewClient(&client)
	if !ok {
		log.Printf("Could not create client: %s", err)
	}
	clientSecret := models.GenerateRandomString(64)
	client.UpdateSecret(clientSecret)
	services.SaveClient(&client)
	localClient := Client{
		Name:   client.Name,
		Key:    client.Key,
		Secret: clientSecret,
		Model:  client,
	}
	return &localClient
}

func NewClientWithScopes(scopes string) *Client {
	client := models.Client{
		Name:         gofakeit.Company(),
		Scopes:       scopes,
		CanonicalURI: []string{ "http://localhost" },
		RedirectURI:  []string{ "http://localhost/callback" },
		Type:         models.ConfidentialClient,
	}
	ok, err := services.CreateNewClient(&client)
	if !ok {
		log.Printf("Could not create client: %s", err)
	}
	clientSecret := models.GenerateRandomString(64)
	client.UpdateSecret(clientSecret)
	services.SaveClient(&client)
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
