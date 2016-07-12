package models

type Session struct {
    Model
    UUID string                 `gorm:"not null;unique;index" validate:"omitempty,uuid4"`
    Owner User                  `gorm:"not null" validate:"exists"`
    Maintainer Client           `gorm:"not null" validate:"exists"`
    Moment int                  `gorm:"not null" validate:"required"`
    Ip string                   `gorm:"not null;index" validate:"required"`
    UserAgent string            `gorm:"not null"`
    Invalidated bool            `gorm:"not null"`
}
