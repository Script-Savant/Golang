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
	userID, exists := c.Get("user_id")
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
