/*
Handle user profile and address management
1. Profile handler: return user info
2. Address handlesrs: crud for user addresses
*/

package controllers

import (
	"ecommerce-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

type AddressInput struct {
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	State     string `json:"state" binding:"required"`
	ZipCode   string `json:"zip_code" binding:"required"`
	Country   string `json:"country" binding:"required"`
	IsDefault bool   `json:"is_default"`
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetProfile(c *gin.Context) {
	// 1. Get user id from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// 2. Find user in database
	var user models.User
	if err := uc.DB.Preload("Address").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 3. Return user profile
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"CreatedAt": user.CreatedAt,
		"addresses": user.Addresses,
	})
}

func (uc *UserController) AddAddress(c *gin.Context) {
	// 1. Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// 2. Bind and validate input
	var input AddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. If setting as default, unset any existing default address
	if input.IsDefault {
		uc.DB.Model(&models.Address{}).
			Where("user_id = ? AND is_default = true", userID).
			Update("is_default", false)
	}

	// 4. Create new address
	address := models.Address{
		UserID:    userID.(uint),
		Street:    input.Street,
		City:      input.City,
		State:     input.State,
		ZipCode:   input.ZipCode,
		Country:   input.Country,
		IsDefault: input.IsDefault,
	}

	if err := uc.DB.Create(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address"})
		return
	}

	// 5. Return created address
	c.JSON(http.StatusCreated, address)
}
