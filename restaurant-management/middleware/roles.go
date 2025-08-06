// allow access to users with specific roles

package middleware

import (
	"golang-restaurant-management/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminOnly -> allow only admin access
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context){
		role := helpers.GetUserRole(c)
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return 
		}
		c.Next()
	}
}