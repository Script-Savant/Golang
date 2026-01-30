package main

import (
	"log"
	"url-shortener/internal/config"
	"url-shortener/internal/models"
	"url-shortener/internal/routes"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	base := "internal/templates/base.html"
	temp_path := "internal/templates/"

	r.AddFromFiles("index", base, temp_path+"index.html")
	r.AddFromFiles("result", base, temp_path+"result.html")

	return r
}

func main() {
	config.ConnectDatabase()
	if err := config.DB.AutoMigrate(&models.URL{}); err != nil {
		log.Fatal("Failed to migrate database")
	}

	r := gin.Default()
	r.HTMLRender = createRenderer()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
