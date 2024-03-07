package main

import (
	"auth/internal/adapters/handler"
	"auth/internal/adapters/repo"
	"auth/internal/core/domain"
	"auth/internal/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	authService *services.AuthService
)

func main() {
	router := fiber.New()
	router.Use(cors.New())
	
	// Connection to PostgreSQL
	dsn := "user=postgres password=testpass123 dbname=go-hex port=3500 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Users{})

	store := repo.NewDB(db)

	authService = services.NewAuthService(store)

	v1 := router.Group("/api")

	authHandler := handler.NewAuthHandler(*authService)
	v1.Post("/signup", authHandler.SignUp)

	router.Listen(":8080")
}