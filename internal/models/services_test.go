package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

func TestValidServiceModel(t *testing.T) {
	service := Service{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         AttachedService,
	}

	assert.True(t, IsValid("validate", service))
}

func TestInvalidServiceMissingRequiredFields(t *testing.T) {
	service := Service{
		Name:         "",
		Description:  "",
		CanonicalURI: "",
		Type:         "",
	}

	assert.False(t, IsValid("validate", service))
}

func TestServiceModelCreation(t *testing.T) {
	var err error

	service := Service{
		Name:         gofakeit.Company(),
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         PublicService,
	}

	assert.True(t, IsValid("validate", service))
	err = service.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	err = service.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%s", err))
}

func TestServiceMarshalJSON(t *testing.T) {
	service := Service{
		Name:         "testing",
		Description:  gofakeit.ProductDescription(),
		CanonicalURI: gofakeit.URL(),
		Type:         PublicService,
	}

	assert.True(t, IsValid("validate", service))
	err := validateModel("validate", service)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = service.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
	err = service.BeforeCreate(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))

	jsonData, err := json.Marshal(service)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "\"name\":\"testing\"")
}
