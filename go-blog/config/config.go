package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	/*
	load environment variables from .env file
	create mysql dsn from env variables
	open db connection with gorm
	assign db connection to DB
	*/

	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	// create a dsn connection form env variables
	user := os.Getenv("DB_USER")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", user, pass, host, port, dbName)

	// open db connection with gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to create a database connection")
	}

	// assign connection to DB
	DB = db
	fmt.Println("Connection to database established")
}