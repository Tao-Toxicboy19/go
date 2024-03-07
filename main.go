package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	
	// Connection to PostgreSQL
	dsn := "user=postgres password=testpass123 dbname=go-hex port=3500 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate()
}
