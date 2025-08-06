// http handlers for user registration, login, validation, hashing and token generation

package controllers

import (
	"golang-restaurant-management/config"
	"golang-restaurant-management/helpers"
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

// Login -> authenticates a user and returns a jwt token
/*
1. bind login json credentials
2. retrieve user by email
3. verify password using bcrypt
4. generate jwt token
5. return token to client
*/
func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 1. bind json
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// 2. get user by email
	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 3. verify pass using bcrypt
	if err := user.VerifyPassword(credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 4. Generate JWT token
	token, err := helpers.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
