package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestSettingRepository__Create() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "authentication",
		Value:    "false",
	}
	err := repository.Create(&setting)
	s.Require().NoError(err)
}

func (s *RepositoryTestSuite) TestSettingRepository__CreateInvalidValue() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "authentication",
		Value:    "10",
	}
	err := repository.Create(&setting)
	s.Require().Error(err)
}

func (s *RepositoryTestSuite) TestSettingRepository__CreateInvalidKey() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "invalid",
		Value:    "10",
	}
	err := repository.Create(&setting)
	s.Require().Error(err)
}

func (s *RepositoryTestSuite) TestSettingRepository__FindOrDefault() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "authentication",
		Value:    "false",
	}
	err := repository.Create(&setting)
	s.Require().NoError(err)
	user = setting.User

	setting = repository.FindOrDefault(user, "notifications", "system-email-notifications", "authentication")
	s.Require().True(setting.IsSavedRecord())
	value, err := setting.DeserializeValue()
	s.Require().NoError(err)
	s.Require().False(value.(bool))

	setting = repository.FindOrDefault(user, "notifications", "system-email-notifications", "email-address")
	s.Require().True(setting.IsNewRecord())
	value, err = setting.DeserializeValue()
	s.Require().NoError(err)
	s.Equal("", value.(string))
	s.Equal("", setting.Value)
	s.Equal(user.ID, setting.User.ID)
}

func (s *RepositoryTestSuite) TestSettingRepository__FindOrFallback() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "authentication",
		Value:    "false",
	}
	err := repository.Create(&setting)
	s.Require().NoError(err)
	user = setting.User

	setting = repository.FindOrFallback(user, "notifications", "system-email-notifications", "authentication", "true")
	s.Require().True(setting.IsSavedRecord())
	value, err := setting.DeserializeValue()
	s.Require().NoError(err)
	s.Require().False(value.(bool))

	setting = repository.FindOrFallback(user, "notifications", "system-email-notifications", "email-address", "example@example.com")
	s.Require().True(setting.IsNewRecord())
	value, err = setting.DeserializeValue()
	s.Require().NoError(err)
	s.Equal("example@example.com", value.(string))
	s.Equal("example@example.com", setting.Value)
	s.Equal(user.ID, setting.User.ID)
}

func (s *RepositoryTestSuite) TestSettingRepository__GetAllForUser() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "authentication",
		Value:    "false",
	}
	err := repository.Create(&setting)
	s.Require().NoError(err)
	user = setting.User

	settings := repository.GetAllForUser(user)
	s.Equal(1, len(settings))
}

func (s *RepositoryTestSuite) TestSettingRepository__ReduceForUser() {
	repository := NewSettingRepository(s.db)

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
	setting := models.Setting{
		User:     user,
		Realm:    "notifications",
		Category: "system-email-notifications",
		Property: "authentication",
		Value:    "false",
	}
	err := repository.Create(&setting)
	s.Require().NoError(err)
	user = setting.User

	reduced := repository.ReduceForUser(user)
	value, ok := reduced["notifications.system-email-notifications.authentication"]
	s.Require().True(ok)
	s.Equal(false, value)
}
