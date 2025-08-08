// GetTables
// GetTable
// CreateTable
// UpdateTable
// DeleteTable

package controllers

import (
	"golang-restaurant-management/config"
	"golang-restaurant-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTables - list all tables
/*
1. Configure pagination
2. Fetch all tables from db
3. Return all tables
*/
func GetTables(c *gin.Context) {
	// 1. pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	// 2. Fetch tables
	var tables []models.Table
	if err := config.DB.Offset(offset).Limit(limit).Find(&tables).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
		"page":   page,
		"limit":  limit,
	})
}
