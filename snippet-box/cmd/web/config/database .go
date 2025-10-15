package config

import (
	"log"

	"github.com/Script-Savant/Golang/snippet-box/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDB(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to make a database connection", err)
	}
	log.Println("Database connection established.")

	if err := db.AutoMigrate(
		&models.Snippet{},
		&models.User{},
		); err != nil {
		log.Fatal("Failed to migrate database tables", err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
