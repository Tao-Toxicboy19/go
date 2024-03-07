package main

import (
	"auth/internal/adapters/repo"
	"auth/internal/core/domain"
	"auth/internal/core/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	authService *services.AuthService
)

func main() {

	// Connection to PostgreSQL
	dsn := "user=postgres password=testpass123 dbname=go-hex port=3500 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Users{})

	store := repo.NewDB(db)

	authService = services.NewAuthService(store)
}
