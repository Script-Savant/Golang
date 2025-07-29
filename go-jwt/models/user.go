package models

import (
	"go-jwt/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func MigrateTables(){
	// create the user table based on the user struct using the connection made in the config pkg
	config.DB.AutoMigrate(&User{})
}