package routes

import (
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/config"
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/handlers"
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	db := config.GetDB()
	authController := handlers.NewAuthController(db)

	router.GET("/register", authController.Register)
	router.POST("/register", authController.Register)
	router.GET("/login", authController.Login)
	router.POST("/login", authController.Login)

	// private routes
	auth := router.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/profile", authController.Profile)
		auth.POST("/profile", authController.Profile)
		auth.GET("/logout", handlers.Logout)
	}
} 