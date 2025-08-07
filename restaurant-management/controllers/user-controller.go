// http handlers for user registration, login, validation, hashing and token generation

package controllers

import (
	"golang-restaurant-management/config"
	"golang-restaurant-management/helpers"
	"golang-restaurant-management/models"
	"net/http"
	"strconv"

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

//Me -> returns the authenticated user's info
/*
1. get user ID from context
2. Find user in db
3. Return user
*/
func Me(c *gin.Context) {
	// 1. get user id from context
	userID := helpers.GetUserID(c)

	// 2. find user in db
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 3. return user
	c.JSON(200, gin.H{"user": user})
}

// UpdateMe -> allows logged in user to update their profile(name and pass)
/*
1. Bind input
2. Get user from db
3. Update only provided fields
4. Save and return updated user
*/
func UpdateMe(c *gin.Context) {
	// 1. bind input
	userID := helpers.GetUserID(c)

	// 2. Get user from DB
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// 3. Only update provided fields
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Password != "" {
		user.Password = input.Password
	}

	// 4. save and return updates user
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(200, gin.H{"message": "Profile updated successfully"})
}

// ListAllUsers -> list all registered users (admin only functionlaity)
/*
1. pagination
2. query all users
3. return the users as JSON
*/
func ListAllUsers(c *gin.Context) {
	// 1. pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	// 2. query all users
	var users []models.User
	if err := config.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch users"})
		return
	}

	// 3. return all users in json
	c.JSON(200, gin.H{
		"users": users,
		"page":  page,
		"limit": limit,
	})
}
