package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/brianvoe/gofakeit/v7"
)

func TestValidServiceModel(t *testing.T) {
	service := Service{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         AttachedService,
	}

	assert.True(t, IsValid("validate", service), "should return true for the valid service")
}

func TestInvalidServiceMissingRequiredFields(t *testing.T) {
	service := Service{
		Name:         "",
		Description:  "",
		CanonicalURI: "",
		Type:         "",
	}

	assert.False(t, IsValid("validate", service), "should return false for missing required fields")
}

func TestServiceModelCreation(t *testing.T) {
	var err error

	service := Service{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         PublicService,
	}

	assert.True(t, IsValid("validate", service), "should return true for the valid service")
	err = service.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	err = service.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}
