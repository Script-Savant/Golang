/*
 open mysql connect, run migrations and seed sample data
1. open gorm db using mysql driver and dsn
2. AutoMigrate models
3. Seed sample products
*/

package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenDatabase opens a GORM DB connection using MySQL DSN from config.
func OpenDatabase(cfg *Config) (*gorm.DB, error) {
	// Steps:
	// 1. Use cfg.DB_DSN to open mysql connection
	// 2. Configure GORM options (if needed)
	// 3. Return *gorm.DB or error
	dsn := cfg.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MigrateAndSeed runs automigrations and seeds sample data.
func MigrateAndSeed(db *gorm.DB) error {
	// Steps:
	// 1. AutoMigrate all models
	// 2. Seed simple products if none exist
	// 3. Return any error
	if err := db.AutoMigrate(&User{}, &Address{}, &Product{}, &Cart{}, &CartItem{}); err != nil {
		return err
	}

	// seed products if table empty
	var count int64
	db.Model(&Product{}).Count(&count)
	if count == 0 {
		log.Println("seeding sample products")
		sample := []Product{
			{SKU: "SKU-1001", Name: "T-Shirt", PriceCents: 1999, Quantity: 100},
			{SKU: "SKU-1002", Name: "Mug", PriceCents: 999, Quantity: 50},
		}
		for _, p := range sample {
			db.Create(&p)
		}
	}
	return nil
}