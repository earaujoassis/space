package repository

import (
	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestLanguageRepository__Create() {
	repository := NewLanguageRepository(s.DB)

	language := models.Language{}
	err := repository.Create(&language)
	s.Error(err)

	language = models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	err = repository.Create(&language)
	s.NoError(err)
}

func (s *RepositoryTestSuite) TestLanguageRepository__GetByID() {
	repository := NewLanguageRepository(s.DB)

	language := models.Language{
		Name:    "Español",
		IsoCode: "es",
	}
	err := repository.Create(&language)
	s.Require().NoError(err)

	retrieved, err := repository.GetByID(language.ID)
	s.Require().NoError(err)
	s.Equal(retrieved.Name, language.Name)
	s.Equal(retrieved.Name, "Español")
}

func (s *RepositoryTestSuite) TestLanguageRepository__GetAllAndCount() {
	repository := NewLanguageRepository(s.DB)

	language := models.Language{
		Name:    "Español",
		IsoCode: "es",
	}
	err := repository.Create(&language)
	s.Require().NoError(err)

	language = models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	err = repository.Create(&language)
	s.Require().NoError(err)

	language = models.Language{
		Name:    "日本語",
		IsoCode: "ja",
	}
	err = repository.Create(&language)
	s.Require().NoError(err)

	language = models.Language{
		Name:    "中文",
		IsoCode: "zh",
	}
	err = repository.Create(&language)
	s.Require().NoError(err)

	all, err := repository.GetAll()
	s.Require().NoError(err)
	s.Equal(len(all), 4)
	s.Equal(all[0].Name, "Español")
	s.Equal(all[1].Name, "Português (Brasil)")
	s.Equal(all[2].Name, "日本語")
	s.Equal(all[3].Name, "中文")

	count, err := repository.Count()
	s.Require().NoError(err)
	s.Equal(int(count), 4)
}

func (s *RepositoryTestSuite) TestLanguageRepository__Save() {
	repository := NewLanguageRepository(s.DB)

	language := models.Language{
		Name:    "日本語",
		IsoCode: "ja",
	}
	err := repository.Create(&language)
	s.Require().NoError(err)
	id := language.ID

	language.Name = "日本語 (nihongo)"
	err = repository.Save(&language)
	s.Require().NoError(err)
	s.Equal(language.Name, "日本語 (nihongo)")
	s.Equal(language.ID, id)
}

func (s *RepositoryTestSuite) TestLanguageRepository__Delete() {
	repository := NewLanguageRepository(s.DB)

	language := models.Language{
		Name:    "中文",
		IsoCode: "zh",
	}
	err := repository.Create(&language)
	s.Require().NoError(err)
	id := language.ID

	err = repository.Delete(id)
	s.Require().NoError(err)
	_, err = repository.GetByID(id)
	s.Error(err)
	s.Equal(err.Error(), "record not found")
}

func (s *RepositoryTestSuite) TestLanguageRepository__FindWhere() {
	repository := NewLanguageRepository(s.DB)

	language := models.Language{
		Name:    "Español",
		IsoCode: "es",
	}
	err := repository.Create(&language)
	s.Require().NoError(err)

	language = models.Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}
	err = repository.Create(&language)
	s.Require().NoError(err)

	language = models.Language{
		Name:    "日本語",
		IsoCode: "ja",
	}
	err = repository.Create(&language)
	s.Require().NoError(err)

	language = models.Language{
		Name:    "中文",
		IsoCode: "zh",
	}
	err = repository.Create(&language)
	s.Require().NoError(err)

	retrieved, err := repository.FindWhere("name = ?", "日本語")
	s.Require().NoError(err)
	s.Equal(len(retrieved), 1)
	s.Equal(retrieved[0].Name, "日本語")
	s.Equal(retrieved[0].IsoCode, "ja")
}
