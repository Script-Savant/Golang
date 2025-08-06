// retrieve user_id and role from gin context

package helpers

import "github.com/gin-gonic/gin"

// get user_id from gin context
func GetUserID(c *gin.Context) uint {
	uid, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return uid.(uint)
}

// get user role from gin context
func GetUserRole(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		return ""
	}
	return role.(string)
}
