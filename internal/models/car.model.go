package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Make               string    `gorm:"size:255;not null"`
	Name               string    `gorm:"size:255;not null"`
	Year               int       `gorm:"not null"`
	PricePerDay        float64   `gorm:"not null"`
	AvailabilityStatus bool      `gorm:"not null"`
}
