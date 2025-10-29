package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi", host, user, pass, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Databse connection opened successfully")

	if migrateErr := db.AutoMigrate(&Book{}); migrateErr != nil {
		log.Fatalf("Failed to migrate database: %v", migrateErr)
	}
	fmt.Println("Database migration complete")

	DB = db
}

func CreateBook(c *gin.Context) {
	var input Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRecord := Book{
		Title:  input.Title,
		Author: input.Author,
	}
	DB.Create(&newRecord)

	c.JSON(http.StatusOK, gin.H{"data": newRecord})
}

func GetBooks(c *gin.Context) {
	var records []Book
	DB.Find(&records)
	c.JSON(http.StatusOK, gin.H{"data": records})
}

func GetBookByID(c *gin.Context) {
	var record Book
	if err := DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record})
}

func UpdateBook(c *gin.Context) {
	var record Book
	if err := DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var input Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.Title = input.Title
	record.Author = input.Author
	DB.Save(&record)

	c.JSON(http.StatusOK, gin.H{"data": record})
}

func DeleteBook(c *gin.Context) {
	var record Book
	if err := DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func main() {
	ConnectDatabase()

	r := gin.Default()

	r.POST("/books", CreateBook)
	r.GET("/books", GetBooks)
	r.GET("/books/:id", GetBookByID)
	r.PATCH("/books/:id", UpdateBook)
	r.DELETE("books/:id", DeleteBook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
