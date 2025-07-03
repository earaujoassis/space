package utils

import (
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/models"
)

func RunUnitTestMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Language{},
		&models.Client{},
		&models.Service{},
		&models.Session{},
		&models.User{},
		&models.Email{},
		&models.Setting{})
}
