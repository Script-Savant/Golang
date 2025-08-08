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

// GetTable - get a specific table details
/*
1. Fetch table id from param
2. Fetch table from db
3. return response with the table
*/
func GetTable(c *gin.Context) {
	// 1. Fetch id from params
	tableID := c.Param("id")

	// 2. Fetch table from db
	var table models.Table
	if err := config.DB.First(&table, tableID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Table fetched successfully",
		"table":   table,
	})
}

// CreateTable - create a table
/*
1. bind json data with table struct
2. Save the table to the db
3. response
*/
func CreateTable(c *gin.Context) {
	// 1. bind incoming json to table struct
	var table models.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Save to db
	if err := config.DB.Create(&table).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "table created successfully", "table": table})
}

// UpdateTable - update a table
/*
1. bind incoming json to a table struct
2. fetch the specified table
3. update the table
4. persist the changes to db
5. return table
*/
func UpdateTable(c *gin.Context) {
	// 1. bindincoming json to table struct
	var body models.Table

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Fetch the candidate table
	tableID := c.Param("id")
	var table models.Table

	if err := config.DB.First(&table, tableID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
		return
	}

	// 3. update the table & persist the changes to db
	if err := config.DB.Model(&table).Updates(&body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Table updated successfully",
		"table":   table,
	})
}

// DeleteTable - delete a specified table
/*
1. fetch table id from params
2. Confirm the table exists in the db
3. delete the table
4. response
*/
func DeleteTable(c *gin.Context) {
	// 1. fwtch id from params
	tableID := c.Param("id")

	// 2. fetch table
	var table models.Table
	if err := config.DB.First(&table, tableID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
		return
	}

	// 3.Delete the table
	if err := config.DB.Delete(&table).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the thable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Table deleted successfully",
	})
}
