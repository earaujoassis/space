package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

type User struct {
	Username      string
	Email         string
	Passphrase    string
	CodeSecretKey string
	Model         models.User
}

func (f *TestRepositoryFactory) NewUser() *User {
	passphrase := gofakeit.Password(true, true, true, true, false, 32)
	user := models.User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: passphrase,
	}
	f.manager.Users().SetRecoverSecret(&user)
	codeSecretKey := f.manager.Users().SetCodeSecret(&user).Secret()
	user.Client = f.manager.Clients().FindOrCreate(models.DefaultClient)
	user.Language = f.manager.Languages().FindOrCreate("English", "en-US")
	f.manager.Users().Create(&user)
	localUser := User{
		Username:      user.Username,
		Email:         user.Email,
		Passphrase:    passphrase,
		CodeSecretKey: codeSecretKey,
		Model:         user,
	}
	return &localUser
}
