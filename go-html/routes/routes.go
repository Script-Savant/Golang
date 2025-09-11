package routes

import (
	"go-html/config"
	"go-html/controllers"
	"go-html/middleware"

	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine) {
	db := config.DB

	r.GET("/", controllers.IndexHandler(db))
	r.GET("/post/:id", controllers.PostHandler(db))

	// Auth routes
	r.GET("/login", controllers.LoginPageHandler)
    r.POST("/login", controllers.LoginHandler(db))
    r.GET("/register", controllers.RegisterPageHandler)
    r.POST("/register", controllers.RegisterHandler(db))
    r.POST("/logout", controllers.LogoutHandler)

	 // Protected routes
    auth := r.Group("/")
    auth.Use(middleware.AuthRequired())
    {
        auth.GET("/create", controllers.CreatePageHandler)
        auth.POST("/create", controllers.CreatePostHandler(db))
        auth.GET("/edit/:id", controllers.EditPageHandler(db))
        auth.POST("/edit/:id", controllers.EditPostHandler(db))
        auth.POST("/delete/:id", controllers.DeletePostHandler(db))
    }
}