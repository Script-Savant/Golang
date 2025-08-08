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
		tables.GET("tables", controllers.GetTables)
		tables.GET("tables/:id", controllers.GetTable)
		tables.POST("tables", middleware.AdminOnly(), controllers.CreateTable)
		tables.PUT("tables/:id", controllers.UpdateTable)
		tables.DELETE("tables/:id", middleware.AdminOnly(), controllers.DeleteTable)
	}
}
