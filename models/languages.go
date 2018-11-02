package models

// Language model/struct represents a Language option through the Application UI
type Language struct {
    Model
    Name string                 `gorm:"not null;unique;index" validate:"required,min=3"`
    IsoCode string              `gorm:"not null;unique" validate:"required,min=2,max=5"`
}
