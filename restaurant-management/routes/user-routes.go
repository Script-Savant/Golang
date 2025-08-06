// routes related to user authentication(register and login)

package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes - add register and login endpoints to the the router
// public routes
func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("api/users/")
	{
		users.POST("register", controllers.Register)
		users.POST("login", controllers.Login)
	}
}
