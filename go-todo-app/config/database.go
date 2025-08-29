package config

import (
	"fmt"
	"go-todo-app/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() *gorm.DB {
	dsn := "alex:password@tcp(localhost:3306)/todo_app?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error creating a db connection")
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("error migrating the database")
	}
	fmt.Println("Connection to the database established successfully")

	DB = db

	return db
}
