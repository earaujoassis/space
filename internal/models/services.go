package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// Service is a service related to this application instance
type Service struct {
	Model
	UUID         string `gorm:"not null;unique;index" validate:"omitempty,uuid4" json:"id"`
	Name         string `gorm:"not null;unique;index" validate:"required,min=3,max=50" json:"name"`
	Description  string `json:"description"`
	CanonicalURI string `gorm:"not null" validate:"required,http_url" json:"uri"`
	LogoURI      string `json:"logo_uri"`
	Type         string `gorm:"not null" validate:"required,service" json:"-"`
}

// BeforeSave Service model/struct hook
func (service *Service) BeforeSave(tx *gorm.DB) error {
	return validateModel("validate", service)
}

// BeforeCreate Service model/struct hook
func (service *Service) BeforeCreate(tx *gorm.DB) error {
	service.UUID = generateUUID()
	return nil
}

func (service Service) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		CanonicalURI string `json:"uri"`
		LogoURI      string `json:"logo_uri"`
	}{
		Id:           service.UUID,
		Name:         service.Name,
		Description:  service.Description,
		CanonicalURI: service.CanonicalURI,
		LogoURI:      service.LogoURI,
	})
}
