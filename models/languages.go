package models

type Language struct {
    Model
    Name string                 `gorm:"not null;unique;index" validate:"required,min=3"`
    IsoCode string              `gorm:"not null;unique" validate:"required,min=2,max=5"`
}
