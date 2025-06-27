package factory

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/internal/models"
)

type Service struct {
	Model models.Service
}

func (f *TestRepositoryFactory) NewService() *Service {
	service := models.Service{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: "http://localhost",
		Type:         models.AttachedService,
	}
	f.manager.Services().Create(&service)
	localService := Service{
		Model: service,
	}
	return &localService
}
