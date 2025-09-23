package main

import (
	"go-stripe/handlers"
	"go-stripe/models"
	"go-stripe/routes"
	"go-stripe/utils"

	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	//Initialize database
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Automigrate models
	if err := db.AutoMigrate(&models.User{}, &models.Transaction{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize stripe
	handlers.InitStripe()

	// setup gin
	r := gin.Default()
	r.HTMLRender = utils.SetupTemplates()

	// session middleware
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	r.Use(sessions.Sessions("mysession", store))

	// setup routes
	routes.SetupRoutes(r, db)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
