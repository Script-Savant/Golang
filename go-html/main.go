package main

import (
	"fmt"
	"go-html/config"
	"go-html/routes"
	"go-html/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/render"
)

func main() {
	// connect database
	config.ConnectDatabase()

	// Iniialize gin
	r := gin.Default()

	// setup sessions
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("blog_session", store))

	// Load templates

	r.HTMLRender = utils.LoadTemplates()

	// serve static files
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")

	// setup routes
	routes.SetupRoutes(r)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
