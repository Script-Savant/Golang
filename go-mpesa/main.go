package main

import (
	"go-mpesa/db"
	"go-mpesa/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init DB
	db.InitDB()

	// Init Gin router
	r := gin.Default()

	// Routes
	r.POST("/c2b", handlers.C2BHandler)           // trigger STK push
	r.POST("/callback/c2b", handlers.C2BCallback) // Safaricom callback

	// Run server
	r.Run(":8080")
}
