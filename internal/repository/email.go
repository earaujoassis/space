package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
)

type EmailRepository struct {
	*BaseRepository[models.Email]
}

func NewEmailRepository(db *database.DatabaseService) *EmailRepository {
	return &EmailRepository{
		BaseRepository: NewBaseRepository[models.Email](db),
	}
}

// GetAllForUser lists emails for a given user
func (r *EmailRepository) GetAllForUser(user models.User) []models.Email {
	emails := make([]models.Email, 0)

	if err := r.db.GetDB().Where("user_id = ?", user.ID).
		Order("address ASC").
		Find(&emails).Error; err != nil {
		return make([]models.Email, 0)
	}

	return emails
}
