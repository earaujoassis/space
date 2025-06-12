package factory

import (
	"log"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
)

type User struct {
	Username      string
	Email         string
	Passphrase    string
	CodeSecretKey string
	Model         models.User
}

func NewUser() *User {
	passphrase := gofakeit.Password(true, true, true, true, false, 32)
	user := models.User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: passphrase,
	}
	user.GenerateRecoverSecret()
	codeSecretKey := user.GenerateCodeSecret().Secret()
	ok, err := services.CreateNewUser(&user)
	if !ok {
		log.Printf("Could not create user: %s", err)
	}
	localUser := User{
		Username:      user.Username,
		Email:         user.Email,
		Passphrase:    passphrase,
		CodeSecretKey: codeSecretKey,
		Model: user,
	}
	return &localUser
}
