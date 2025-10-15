package utils

import (
	"github.com/Script-Savant/Golang/snippet-box/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/gin-contrib/sessions"
)

func GetCurrentUser(c *gin.Context, db *gorm.DB) *models.User {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID != nil {
		var user models.User
		db.First(&user, userID)
		return &user
	}
	return nil
}