package controllers

import (
	"go-blog/models"
	"go-blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthController handles authentication related operations
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController -> create a new AuthController with db connection
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

// Handle user registration
func (ac *AuthController) Register(c *gin.Context) {
	/*
		- Bind the incoming json to a user struct
		- Hash the password before storing it
		- create the user in the database
		- create a default profile for the user
		- Return success response
	*/

	// 1. bind the icoming json to the user struct
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	// 3. create the user in the database
	if err := ac.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// 4. create a default profile for the user
	profile := models.Profile{
		UserID: user.ID,
		Bio:    "New user",
	}
	if err := ac.DB.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating profile"})
		return
	}

	// 5. Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login handles user login and jwt token generation
func (ac *AuthController) Login(c *gin.Context) {
	/*
		- bind the icoming json to a login request struct
		- find the user by email
		- compare the provided password with the stored hash
		- generate jwt token
		- return the token in the response
	*/

	// 1. bind the incoming json to login request struct
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. find the user by email
	var user models.User
	if err := ac.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 3. compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 4. generate jwt token
	jwtWrapper := utils.NewJWTWrapper()
	token, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// 5. return the token in the response
	c.JSON(http.StatusOK, gin.H {
		"token": token,
		"message": "Successfully logged in",
	})
}
