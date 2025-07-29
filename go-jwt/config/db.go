package config

import(
	"log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // hold the database connection

func ConnectDatabase() {
	/*
	Connect to mysql db
	open the connection using gorm
	store the connection to the DB variable
	*/
	dsn := "root:password@tcp(127.0.0.1:3306)/go_jwt_auth"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	fmt.Println("Database connection established...")
}