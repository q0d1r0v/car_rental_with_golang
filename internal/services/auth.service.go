package services

import (
	"car_rental_with_golang/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// Register user (user yaratish va parolni shifrlash)
func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	// Foydalanuvchi mavjudligini tekshirish
	var existingUser models.User
	if err := s.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		// Agar foydalanuvchi mavjud bo'lsa, xatolik qaytarish
		return nil, fmt.Errorf("username %s already taken", username)
	}

	if err := s.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		// Agar email mavjud bo'lsa, xatolik qaytarish
		return nil, fmt.Errorf("email %s already taken", email)
	}

	// Parolni shifrlash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Yangi foydalanuvchini yaratish
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Foydalanuvchini saqlash
	result := s.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// Login user (foydalanuvchini tekshirish va JWT tokenini yaratish)
func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	// Email bo'yicha foydalanuvchini qidirish
	result := s.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("user not found")
	}

	// Parolni tekshirish
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	// JWT token yaratish
	token, err := s.generateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// JWT yaratish
func (s *AuthService) generateJWT(userID uuid.UUID, username string, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userID.String(),                       // userID ni string sifatida saqlash
		"username": username,                              // username qo'shish
		"email":    email,                                 // email qo'shish
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token 24 soat davomida amal qiladi
		"iat":      time.Now().Unix(),                     // Token yaratish vaqti
	}

	// JWTni yaratish
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tokenni imzolash
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign the token: %v", err)
	}

	return signedToken, nil
}
