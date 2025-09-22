package db

import (
	"go-mpesa/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("go_mpesa.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to make a  database connection")
	}
	
	if err := db.AutoMigrate(
		&models.Transaction{},
	); err != nil {
		log.Fatal("Failed to migrate tables")
	}

	DB = db
}

func GetDB() *gorm.DB{
	return DB
}