package factory

import (
	"github.com/earaujoassis/space/internal/models"
)

type Group struct {
	Model models.Group
}

func (f *TestRepositoryFactory) NewGroup(user models.User, client models.Client) *Group {
	session := models.Group{
		User:   user,
		Client: client,
		Tags:   []string{"testing"},
	}
	f.manager.Groups().Create(&session)
	localGroup := Group{
		Model: session,
	}
	return &localGroup
}
