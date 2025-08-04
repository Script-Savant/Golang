package main

import (
	"go-blog/config"
	"go-blog/models"
	"go-blog/routes"
	"log"
)

func main() {
	// connect database
	config.ConnectDatabase()

	// auto migrate models
	db := config.DB
	err := db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Post{},
		&models.Tag{},
		&models.Comment{},
		&models.Like{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// setup the router with all routes
	router := routes.SetupRoutes(db)

	// start the server
	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
