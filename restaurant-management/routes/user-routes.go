// routes related to user authentication(register and login)

package routes

import (
	"golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes - add register and login endpoints to the the router
func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("api/users/")
	{
		// public routes
		users.POST("register", controllers.Register)
		users.POST("login", controllers.Login)

		// protected routes
		users.GET("me", middleware.JWTAuthMiddleware(), controllers.Me)
		users.PUT("me", middleware.JWTAuthMiddleware(), controllers.UpdateMe)
		users.GET("all-users", middleware.JWTAuthMiddleware(), middleware.AdminOnly(), controllers.ListAllUsers)
		users.DELETE("delete-user", middleware.JWTAuthMiddleware(), middleware.AdminOnly(), controllers.DeleteUser)
	}
}
