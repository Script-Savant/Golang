package main

/*
- entry point for the restaurant mgmt api
- initialize db connection, set up routes, start the server
*/

import (
	"fmt"
	"golang-restaurant-management/config"
	"golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize db connection
	config.SetupDatabase()
	fmt.Println("Database connected successfully")

	// create gin router
	r := gin.Default()

	// register user routes (public)
	routes.RegisterUserRoutes(r)

	// start server
	fmt.Println("Server starting at localhost:8080")
	r.Run(":8080")
}
