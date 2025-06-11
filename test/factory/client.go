package factory

import (
	"log"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
)

type Client struct {
	Name         string
	Key          string
	Secret       string
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
		Name:          client.Name,
		Key:           client.Key,
		Secret:        clientSecret,
	}
	return &localClient
}
