package config

import (
	"fmt"
	"go-html/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to make databace connection")
	}

	fmt.Println("Database connection established")

	if err := db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Fatal("Database migration failed")
	}
	
	DB = db
}
