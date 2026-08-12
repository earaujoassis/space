package utils

import (
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/models"
)

func RunUnitTestMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Language{},
		&models.Client{},
		&models.Service{},
		&models.User{},
		&models.Session{},
		&models.Email{},
		&models.Setting{},
		&models.Group{},
	)
}
