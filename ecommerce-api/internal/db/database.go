/*
Initialize db connection and auto migrate models
1. Connect to db using GORM
2. Auto-migrate all models
3. Return DB instance
*/

package db

import (
	"ecommerce-api/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initializeDB(cfg *config.Config) (*gorm.DB, error) {
	// 1. Create DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// 2. open database connection
	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if  err != nil {
		return  nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// 3. Auto-migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate models: %w", err)
	}

	log.Println("Database connection established and models migrated")
	return DB, nil
}

func GetDB() *gorm.DB {
	return DB
}