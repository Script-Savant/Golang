package main

import (
	"go-htmx-clients/config"
	"go-htmx-clients/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// db setup
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Client{})

	// routes
	r.GET("/", showClients)
	r.POST("/add", addClient)
	r.POST("/delete/:id", deleteClient)

	r.Run(":8080")
}

func showClients(c *gin.Context) {
	var clients []models.Client
	config.DB.Find(&clients)
	c.HTML(http.StatusOK, "index.html", gin.H{"clients": clients})
}

func addClient(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBind(&client); err == nil {
		config.DB.Create(&client)
	}

	var clients []models.Client
	config.DB.Find(&clients)

	c.HTML(http.StatusOK, "clients.html", gin.H{"clients": clients})
}

func deleteClient(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Client{}, id)

	var clients []models.Client
	config.DB.Find(&clients)

	c.HTML(http.StatusOK, "clients.html", gin.H{"clients": clients})
}