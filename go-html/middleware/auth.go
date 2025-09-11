package middleware

import (
	"go-html/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetCurrentUser(c *gin.Context, db *gorm.DB) *models.User {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		return nil
	}

	var user models.User
	if db != nil {
		db.First(&user, userID)
	} else {
		// For cases where we only need to check if user is logged in
		// but don't need full user details
		user.ID = userID.(uint)
	}
	return &user
}
