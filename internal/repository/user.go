package repository

import (
	"fmt"
	"strings"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"

	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
)

type UserRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository(db *database.DatabaseService) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[models.User](db),
	}
}

// FindByAccountHolder gets an user by its account holder (username or email)
func (r *UserRepository) FindByAccountHolder(holder string) models.User {
	var user models.User
	r.db.GetDB().Preload("Client").Preload("Language").Where("username = ? OR email = ?", holder, holder).First(&user)
	return user
}

// FindByPublicID gets an user by its public ID (used by client applications)
func (r *UserRepository) FindByPublicID(publicID string) models.User {
	var user models.User
	r.db.GetDB().Preload("Client").Preload("Language").Where("public_id = ?", publicID).First(&user)
	return user
}

// FindByUUID gets an user by its UUID (internal use only)
func (r *UserRepository) FindByUUID(uuid string) models.User {
	var user models.User
	r.db.GetDB().Preload("Client").Preload("Language").Where("uuid = ?", uuid).First(&user)
	return user
}

// FindByID gets an user by its ID (internal use only)
func (r *UserRepository) FindByID(id uint) models.User {
	var user models.User
	r.db.GetDB().Preload("Client").Preload("Language").Where("id = ?", id).First(&user)
	return user
}

// ActiveClients lists client applications for a given user
func (r *UserRepository) ActiveClients(user models.User) []models.Client {
	clients := make([]models.Client, 0)

	r.db.GetDB().
		Raw("SELECT DISTINCT clients.uuid, clients.name, clients.description, clients.canonical_uri "+
			"FROM clients JOIN sessions ON clients.id = sessions.client_id "+
			"WHERE sessions.token_type IN ('access_token', 'refresh_token') AND sessions.invalidated = false AND "+
			"sessions.user_id = ?;", user.ID).
		Scan(&clients)
	return clients
}

// Authentic checks if a password + passcode combination is valid for a given User
func (r *UserRepository) Authentic(user models.User, password, passcode string) bool {
	var validPasscode bool
	validPassword := bcrypt.CompareHashAndPassword([]byte(user.Passphrase), []byte(password)) == nil
	codeSecret, err := security.Decrypt(r.db.GetStorageSecret(), user.CodeSecret)
	if err != nil {
		return false
	}

	validPasscode = totp.Validate(passcode, string(codeSecret))
	return validPasscode && validPassword
}

// SetCodeSecret sets a code secret for an user, to generate TOPT codes, without saving
func (r *UserRepository) SetCodeSecret(user *models.User) *otp.Key {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "quatrolabs.com",
		AccountName: user.Username,
	})
	codeSecret := key.Secret()
	if cryptedCodeSecret, err := security.Encrypt(r.db.GetStorageSecret(), []byte(codeSecret)); err == nil {
		user.CodeSecret = string(cryptedCodeSecret)
	} else {
		user.CodeSecret = codeSecret
	}
	if err != nil {
		return nil
	}
	return key
}

// SetRecoverSecret sets a recover secret string for a user without saving
func (r *UserRepository) SetRecoverSecret(user *models.User) (string, error) {
	var secret = strings.ToUpper(fmt.Sprintf("%s-%s-%s-%s",
		models.GenerateRandomString(4),
		models.GenerateRandomString(4),
		models.GenerateRandomString(4),
		models.GenerateRandomString(4)))
	if cryptedRecoverSecret, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost); err == nil {
		user.RecoverSecret = string(cryptedRecoverSecret)
	} else {
		return secret, err
	}
	return secret, nil
}

// SetPassword sets a User's password without saving
func (r *UserRepository) SetPassword(user *models.User, password string) error {
	user.Passphrase = password
	if !models.IsValid("essential", user) {
		return fmt.Errorf("user validation failed")
	}
	crypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		user.Passphrase = string(crypted)
		return nil
	}
	return err
}
