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

	user := r.Group("/users")
	{
		// Auth routes
		user.GET("/login", controllers.LoginPageHandler)
		user.POST("/login", controllers.LoginHandler(db))
		user.GET("/register", controllers.RegisterPageHandler)
		user.POST("/register", controllers.RegisterHandler(db))
		user.POST("/logout", controllers.LogoutHandler)
	}

	posts := r.Group("/posts")
	{
		posts.GET("/", controllers.IndexHandler(db))
		posts.GET("/post/:id", controllers.PostHandler(db))

		// Protected routes
		posts.Use(middleware.AuthRequired())
		{
			posts.GET("/create", controllers.CreatePageHandler)
			posts.POST("/create", controllers.CreatePostHandler(db))
			posts.GET("/edit/:id", controllers.EditPageHandler(db))
			posts.POST("/edit/:id", controllers.EditPostHandler(db))
			posts.POST("/delete/:id", controllers.DeletePostHandler(db))
		}
	}
}
