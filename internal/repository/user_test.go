package repository

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/pquerna/otp/totp"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestUserRepository__FindByAccountHolder() {
	repository := NewUserRepository(s.DB)

	username := gofakeit.Username()
	email := gofakeit.Email()
	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      username,
		Email:         email,
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err := repository.Create(&user)
	s.Require().NoError(err)

	retrievedByEmail := repository.FindByAccountHolder(email)
	s.Require().NotZero(retrievedByEmail.ID)
	s.Equal(retrievedByEmail.Email, email)
	s.Equal(retrievedByEmail.Username, username)

	retrievedByUsername := repository.FindByAccountHolder(username)
	s.Require().NotZero(retrievedByUsername.ID)
	s.Equal(retrievedByUsername.Email, email)
	s.Equal(retrievedByUsername.Username, username)

	anotherEmail := gofakeit.Email()
	unknownUser := repository.FindByAccountHolder(anotherEmail)
	s.Require().Zero(unknownUser.ID)

	anotherUsername := gofakeit.Username()
	unknownUser = repository.FindByAccountHolder(anotherUsername)
	s.Require().Zero(unknownUser.ID)
}

func (s *RepositoryTestSuite) TestUserRepository__FindByPublicID() {
	repository := NewUserRepository(s.DB)

	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err := repository.Create(&user)
	s.Require().NoError(err)
	publicID := user.PublicID

	retrieved := repository.FindByPublicID(publicID)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.Email, user.Email)
	s.Equal(retrieved.Username, user.Username)
}

func (s *RepositoryTestSuite) TestUserRepository__FindByUUID() {
	repository := NewUserRepository(s.DB)

	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err := repository.Create(&user)
	s.Require().NoError(err)
	uuid := user.UUID

	retrieved := repository.FindByUUID(uuid)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.Email, user.Email)
	s.Equal(retrieved.Username, user.Username)
}

func (s *RepositoryTestSuite) TestUserRepository__FindByID() {
	repository := NewUserRepository(s.DB)

	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err := repository.Create(&user)
	s.Require().NoError(err)
	id := user.ID

	retrieved := repository.FindByID(id)
	s.Require().NotZero(retrieved.ID)
	s.Equal(retrieved.ID, id)
	s.Equal(retrieved.Email, user.Email)
	s.Equal(retrieved.Username, user.Username)
}

func (s *RepositoryTestSuite) TestUserRepository__ActiveClients() {
	repository := NewUserRepository(s.DB)

	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err := repository.Create(&user)
	s.Require().NoError(err)

	activeClients := repository.ActiveClients(user)
	s.Equal(len(activeClients), 0)
}

func (s *RepositoryTestSuite) TestUserRepository__Authentic() {
	repository := NewUserRepository(s.DB)

	passphrase := gofakeit.Password(true, true, true, true, false, 32)
	email := gofakeit.Email()
	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         email,
		Passphrase:    passphrase,
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	repository.SetRecoverSecret(&user)
	codeSecretKey := repository.SetCodeSecret(&user).Secret()
	err := repository.Create(&user)
	s.Require().NoError(err)

	code, _ := totp.GenerateCode(codeSecretKey, time.Now())
	fakeCode := "000000"
	fakePassphrase := gofakeit.Password(true, true, true, true, false, 32)

	retrievedByEmail := repository.FindByAccountHolder(email)
	result := repository.Authentic(retrievedByEmail, passphrase, code)
	s.True(result)

	result = repository.Authentic(retrievedByEmail, fakePassphrase, code)
	s.False(result)

	result = repository.Authentic(retrievedByEmail, passphrase, fakeCode)
	s.False(result)

	newPassphrase := gofakeit.Password(true, true, true, true, false, 32)
	err = repository.SetPassword(&user, newPassphrase)
	s.Require().NoError(err)
	err = repository.Save(&user)
	s.Require().NoError(err)

	code, _ = totp.GenerateCode(codeSecretKey, time.Now())
	retrievedByEmail = repository.FindByAccountHolder(email)

	result = repository.Authentic(retrievedByEmail, passphrase, code)
	s.False(result)

	result = repository.Authentic(retrievedByEmail, newPassphrase, code)
	s.True(result)
}
