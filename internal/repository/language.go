package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
)

type LanguageRepository struct {
	*BaseRepository[models.Language]
}

func NewLanguageRepository(db *database.DatabaseService) *LanguageRepository {
	return &LanguageRepository{
		BaseRepository: NewBaseRepository[models.Language](db),
	}
}

// FindOrCreate attempts to obtain a Language entry through
// its `name` and ISO Code; if that's not available, it creates one
func (r *LanguageRepository) FindOrCreate(name, isoCode string) models.Language {
	var language models.Language

	r.db.GetDB().Where("name = ? AND iso_code = ?", name, isoCode).First(&language)
	if language.IsNewRecord() {
		language = models.Language{Name: name, IsoCode: isoCode}
		r.db.GetDB().Create(&language)
	}

	return language
}
