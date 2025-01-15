package services

import (
	"car_rental_with_golang/internal/models"

	"gorm.io/gorm"
)

type RoleService struct {
	DB *gorm.DB
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	result := s.DB.Find(&roles)
	return roles, result.Error
}
