package controllers

import (
	"car_rental_with_golang/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func (c *AuthController) Register(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.AuthService.Register(input.Username, input.Email, input.Password)
	if err != nil {
		log.Println("Error registering user:", err)
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
func (ac *AuthController) Login(c *gin.Context) {
	// Login uchun kerakli malumotlarni olish
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Foydalanuvchi login qilish
	token, err := ac.AuthService.Login(loginData.Email, loginData.Password)
	if err != nil {
		// Xato bo'lsa
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Tokenni muvaffaqiyatli yaratgan taqdirda
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
