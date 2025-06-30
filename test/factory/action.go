package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

type Action struct {
	Model models.Action
}

func (f *TestRepositoryFactory) NewAction(user models.User) *Action {
	client := f.DefaultClient().Model
	action := models.Action{
		User:        user,
		Client:      client,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      models.WriteScope,
		Description: models.UpdateUserAction,
	}
	f.manager.Actions().Create(&action)
	localAction := Action{
		Model: action,
	}
	return &localAction
}

func (f *TestRepositoryFactory) NewActionWithoutPermissions(user models.User) *Action {
	client := f.DefaultClient().Model
	action := models.Action{
		User:        user,
		Client:      client,
		IP:          gofakeit.IPv4Address(),
		UserAgent:   gofakeit.UserAgent(),
		Scopes:      models.ReadScope,
		Description: models.NotSpecialAction,
	}
	f.manager.Actions().Create(&action)
	localAction := Action{
		Model: action,
	}
	return &localAction
}
