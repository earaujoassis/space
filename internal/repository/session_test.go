package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestSessionRepository__FindByUUID() {
	repository := NewSessionRepository(s.DB)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	err := repository.Create(&session)
	s.Require().NoError(err)

	uuid := session.UUID
	retrieved := repository.FindByUUID(uuid)
	s.Require().NotZero(retrieved.ID)
	s.Equal(client.Name, session.Client.Name)
	s.Equal(user.Username, session.User.Username)
	s.Equal(uuid, session.UUID)
}

func (s *RepositoryTestSuite) TestSessionRepository__FindByToken() {
	repository := NewSessionRepository(s.DB)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	err := repository.Create(&session)
	s.Require().NoError(err)

	token := session.Token
	retrieved := repository.FindByToken(token, models.ApplicationToken)
	s.Require().NotZero(retrieved.ID)
	s.Equal(client.Name, session.Client.Name)
	s.Equal(user.Username, session.User.Username)
	s.Equal(token, session.Token)

	retrieved = repository.FindByToken(token, models.AccessToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.RefreshToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.GrantToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.IDToken)
	s.Require().Zero(retrieved.ID)
}

func (s *RepositoryTestSuite) TestSessionRepository__Invalidate() {
	repository := NewSessionRepository(s.DB)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	err := repository.Create(&session)
	s.Require().NoError(err)

	token := session.Token
	repository.Invalidate(&session)

	retrieved := repository.FindByToken(token, models.ApplicationToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.AccessToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.RefreshToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.GrantToken)
	s.Require().Zero(retrieved.ID)

	retrieved = repository.FindByToken(token, models.IDToken)
	s.Require().Zero(retrieved.ID)
}

func (s *RepositoryTestSuite) TestSessionRepository__ApplicationSessions() {
	repository := NewSessionRepository(s.DB)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	err := repository.Create(&session)
	s.Require().NoError(err)

	user = session.User
	sessions := repository.ApplicationSessions(user)
	s.Require().Equal(1, len(sessions))
	s.False(sessions[0].Current)
	sessions = repository.ApplicationSessionsWithActive(user, session)
	s.Require().Equal(1, len(sessions))
	s.True(sessions[0].Current)

	session = models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.GrantToken,
	}
	err = repository.Create(&session)
	s.Require().NoError(err)

	sessions = repository.ApplicationSessions(user)
	s.Require().Equal(1, len(sessions))
}

func (s *RepositoryTestSuite) TestSessionRepository__ActiveForClient() {
	repository := NewSessionRepository(s.DB)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	err := repository.Create(&session)
	s.Require().NoError(err)

	client = session.Client
	user = session.User
	count := repository.ActiveForClient(client, user)
	s.Require().Zero(count)

	session = models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.AccessToken,
	}
	err = repository.Create(&session)
	s.Require().NoError(err)
	session = models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.RefreshToken,
	}
	err = repository.Create(&session)
	s.Require().NoError(err)

	count = repository.ActiveForClient(client, user)
	s.Require().Equal(2, int(count))
}

func (s *RepositoryTestSuite) TestSessionRepository__RevokeAccess() {
	repository := NewSessionRepository(s.DB)

	client := models.Client{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: []string{"http://localhost"},
		RedirectURI:  []string{"http://localhost/callback"},
		Scopes:       models.PublicScope,
		Type:         models.ConfidentialClient,
	}
	language := models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	user := models.User{
		Client:        client,
		Language:      language,
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.ApplicationToken,
	}
	err := repository.Create(&session)
	s.Require().NoError(err)

	client = session.Client
	user = session.User
	count := repository.ActiveForClient(client, user)
	s.Require().Zero(count)

	session = models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.AccessToken,
	}
	err = repository.Create(&session)
	s.Require().NoError(err)
	session = models.Session{
		User:      user,
		Client:    client,
		IP:        gofakeit.IPv4Address(),
		UserAgent: gofakeit.UserAgent(),
		Scopes:    models.PublicScope,
		TokenType: models.RefreshToken,
	}
	err = repository.Create(&session)
	s.Require().NoError(err)

	count = repository.ActiveForClient(client, user)
	s.Require().Equal(2, int(count))

	repository.RevokeAccess(client, user)
	count = repository.ActiveForClient(client, user)
	s.Require().Zero(count)
}
