/*
- database configuration and connection setup
- handle mysql database connection using gorm
*/

package config

import (
	"fmt"
	"golang-restaurant-management/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDatabase - creates a connection to mysql database and returns a GORM db instance
/*
Steps
1. Get database credentials from environment variables
2. create a dsn string
3. open a connection using gorm
4. Auto migrate models
*/
func SetupDatabase() {
	// get db credentials from env variables - user, pass, host, port, db name
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, shall use defaults")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("PASS")
	if dbPass == "" {
		dbPass = "password"
	}

	dbHost := os.Getenv("HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "restaurant_api"
	}

	// create a dsn - data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	log.Println("Successfully connected to the database")
	DB = db

	// Automigrate models
	db.AutoMigrate(
		&models.User{},
		&models.Table{},
		&models.Food{},
		&models.Menu{},
		&models.Order{},
		&models.Order{},
		&models.OrderItem{},
		&models.Invoice{},
		&models.Note{},
	)
}
