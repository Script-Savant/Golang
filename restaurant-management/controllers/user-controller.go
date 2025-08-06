// http handlers for user registration, login, validation, hashing and token generation

package controllers

import (
	"golang-restaurant-management/config"
	"golang-restaurant-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register - handle user registration
/*
1. Bind incoming json to a user struct
2. Validate required fields
3. Create user in DB -> BeforeSave will hash password
4. Return success or error
*/
func Register(c *gin.Context) {
	var user models.User

	// 1. bind json to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// 2. basic validation
	if user.Email == "" || user.Password == "" || user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, email and password are required"})
		return
	}

	// 3. create a user in db
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}
