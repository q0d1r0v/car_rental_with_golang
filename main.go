package main

import (
	"car_rental_with_golang/internal/controllers"
	"car_rental_with_golang/internal/repositories"
	"car_rental_with_golang/internal/routes"
	"car_rental_with_golang/internal/services"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// PostgreSQL ulanish
	dsn := "user=postgres password=3801 dbname=mydatabascar_rental_db port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database")
		return
	}

	// Repositories va Service'larni yaratish
	userRepo := &repositories.UserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userController := &controllers.UserController{Service: userService}

	// Gin routerini yaratish
	r := gin.Default()

	// Route'larni sozlash
	routes.SetupRouter(r, userController)

	// Serverni ishga tushurish
	r.Run(":3000")
}
