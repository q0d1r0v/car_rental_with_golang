package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BookingID     uint      `gorm:"not null"`
	Booking       Booking   `gorm:"foreignKey:BookingID"`
	AmountPaid    float64   `gorm:"not null"`
	PaymentDate   time.Time `gorm:"not null"`
	PaymentStatus string    `gorm:"size:50;not null;default:'completed'"`
}
