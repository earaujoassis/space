package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/pquerna/otp/totp"

	"github.com/earaujoassis/space/internal/models"
)

type User struct {
	Passphrase    string
	CodeSecretKey string
	Model         models.User
}

type UserOptions struct {
	Admin bool
}

func (user *User) GenerateCode() string {
	code, _ := totp.GenerateCode(user.CodeSecretKey, time.Now())
	return code
}

func (f *TestRepositoryFactory) GetAvailableUser() models.User {
	if entities, err := f.manager.Users().GetAll(); err == nil {
		id := entities[0].ID
		return f.manager.Users().FindByID(id)
	}

	return models.User{}
}

func (f *TestRepositoryFactory) NewUserWithOption(opts UserOptions) *User {
	repositories := f.manager
	passphrase := gofakeit.Password(true, true, true, true, false, 32)
	user := models.User{
		FirstName:  gofakeit.FirstName(),
		LastName:   gofakeit.LastName(),
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		Passphrase: passphrase,
		Admin:      opts.Admin,
	}
	repositories.Users().SetRecoverSecret(&user)
	codeSecretKey := repositories.Users().SetCodeSecret(&user).Secret()
	user.Client = f.DefaultClient().Model
	user.Language = repositories.Languages().FindOrCreate("English", "en-US")
	repositories.Users().Create(&user)
	localUser := User{
		Passphrase:    passphrase,
		CodeSecretKey: codeSecretKey,
		Model:         user,
	}
	return &localUser
}

func (f *TestRepositoryFactory) NewUser() *User {
	return f.NewUserWithOption(UserOptions{Admin: true})
}
