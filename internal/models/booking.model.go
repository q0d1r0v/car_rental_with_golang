package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	CarID     uint      `gorm:"not null"`
	Car       Car       `gorm:"foreignKey:CarID"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	TotalCost float64   `gorm:"not null;default:0"`
	Status    string    `gorm:"size:50;not null;default:'pending'"`
}
