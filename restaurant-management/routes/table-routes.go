// configure table routes

package routes

import (
	"golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterTableRoutes - routes associated with table functionlaity
func RegisterTableRoutes(r *gin.Engine) {
	tables := r.Group("api/tables/")
	tables.Use(middleware.JWTAuthMiddleware())
	{
		tables.GET("/", controllers.GetTables)
		tables.GET("/:id", controllers.GetTable)
		tables.POST("/", middleware.AdminOnly(), controllers.CreateTable)
		tables.PUT("/:id", controllers.UpdateTable)
		tables.DELETE("/:id", middleware.AdminOnly(), controllers.DeleteTable)
	}
}
