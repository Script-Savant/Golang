package main

import (
	"go-jwt/config"
	"go-jwt/models"
	"go-jwt/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	// connect to mysql db
	config.ConnectDatabase()

	// auto create tables if they do not exist
	models.MigrateTables()

	// register all route groups
	routes.RegisterAuthRoutes(r)

	// start the server
	r.Run(":8080")
}