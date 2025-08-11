// Register all routes
package routes

import (
	"Ecommerce-Cart/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes - create a catalogue of all routes with their corresponding controllers
func UserRoutes(router *gin.Engine) {
	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)
	router.POST("/admin/add-product", controllers.AddProduct)
	router.GET("/users/products-view", controllers.SearchProduct)
	router.GET("/users/search", controllers.SearchProductByQuery)
}