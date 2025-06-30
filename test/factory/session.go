package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

type Session struct {
	Model models.Session
}

func (f *TestRepositoryFactory) NewGrantToken(user models.User) *Session {
	client := f.DefaultClient().Model
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.GrantToken,
	}
	f.manager.Sessions().Create(&session)
	localSession := Session{
		Model: session,
	}
	return &localSession
}

func (f *TestRepositoryFactory) NewApplicationSession(user models.User) *Session {
	client := f.DefaultClient().Model
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	f.manager.Sessions().Create(&session)
	localSession := Session{
		Model: session,
	}
	return &localSession
}

func (f *TestRepositoryFactory) NewRefreshToken(user models.User, client models.Client) *Session {
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.RefreshToken,
	}
	f.manager.Sessions().Create(&session)
	localSession := Session{
		Model: session,
	}
	return &localSession
}
