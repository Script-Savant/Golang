package routes

import (
	"go-jwt/controllers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	// group all auth-related routes under "/auth"
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)

		// Protected route
		auth.GET("/profile", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized to view this profile!"})
		})
	}
}
