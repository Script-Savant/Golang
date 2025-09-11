package main

import (
	"fmt"
	"go-html/config"
	"go-html/routes"
	"html/template"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
	r.SetFuncMap(template.FuncMap{
        "formatDate": func(t time.Time) string {
            return t.Format("Jan 2, 2006")
        },
    })
	r.LoadHTMLGlob("templates/*")

	// serve static files
	r.Static("/static", "./static")
	r.Static("/uploads", "./uploads")

	// setup routes
	routes.SetupRoutes(r)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
