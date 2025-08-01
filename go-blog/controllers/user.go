/*
make profile operations
- create a database connection
- retrieve a profile
- update a profile
*/
package controllers

import (
	"go-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController -> handle user related operations
type UserController struct {
	DB *gorm.DB
}

// NewUserController -> creates a UserController with a database connection
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// Get profile of the current user
func (uc *UserController) GetProfile(c *gin.Context) {
	/*
		- get the email from the context (set by AuthMiddleware)
		- find the user by email
		- return the user profile
	*/
	// 1. get the email from the context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	// 2. find the user by email
	var user models.User
	if err := uc.DB.Preload("Profile").Where("email = ?", email.(string)).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 3. return the user profile
	c.JSON(http.StatusOK, gin.H{
		"profile": user.Profile,
	})
}
