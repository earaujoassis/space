package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

type Email struct {
	Model models.Email
}

func (f *TestRepositoryFactory) NewEmailFor(user models.User) *Email {
	email := models.Email{
		User:    user,
		Address: gofakeit.Email(),
	}
	f.manager.Emails().Create(&email)
	localEmail := Email{
		Model: email,
	}
	return &localEmail
}
