package services

import (
	"car_rental_with_golang/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// Barcha foydalanuvchilarni olish
func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := s.DB.Find(&users)
	return users, result.Error
}
