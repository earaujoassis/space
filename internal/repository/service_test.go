package repository

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestServiceRepository__GetAll() {
	repository := NewServiceRepository(s.DB)

	service := models.Service{
		Name:         "Omega",
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         models.AttachedService,
	}
	err := repository.Create(&service)
	s.Require().NoError(err)

	service = models.Service{
		Name:         "Alpha",
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         models.AttachedService,
	}
	err = repository.Create(&service)
	s.Require().NoError(err)

	retrieved := repository.GetAll()
	s.Equal(len(retrieved), 2)
	s.Equal(retrieved[0].Name, "Alpha")
	s.Equal(retrieved[1].Name, "Omega")
}
