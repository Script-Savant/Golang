package controllers

import (
	"fmt"
	"go-html/middleware"
	"go-html/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func IndexHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(">>> IndexHandler CALLED")
		var posts []models.Post

		if err := db.Preload("Author").Order("created_at desc").Find(&posts).Error; err != nil {
			c.HTML(http.StatusInternalServerError, "index", gin.H{"error": "Failed to fetch posts"})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts, "user": middleware.GetCurrentUser(c, db)})
	}
}

func PostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("id")

		var post models.Post

		if err := db.Preload("Author").First(&post, postID).Error; err != nil {
			c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "post not found"})
			return
		}

		c.HTML(http.StatusOK, "post.html", gin.H{"post": post, "user": middleware.GetCurrentUser(c, db)})
	}
}

func CreatePageHandler(c *gin.Context) {
	user := middleware.GetCurrentUser(c, nil)
	if user == nil {
		c.Redirect(http.StatusFound, "/users/login")
		return
	}

	c.HTML(http.StatusOK, "create.html", gin.H{
		"user": user,
	})
}

func CreatePostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetCurrentUser(c, db)
		if user == nil {
			c.Redirect(http.StatusFound, "/users/login")
			return
		}

		title := c.PostForm("title")
		content := c.PostForm("content")

		var imagePath string

		file, header, err := c.Request.FormFile("image")
		if err == nil {
			defer file.Close()

			if !isImage(header) {
				// Return error for invalid file type
				c.HTML(http.StatusBadRequest, "create.html", gin.H{
					"error": "Please upload a valid image file (JPEG, PNG, GIF, or WebP)",
					"user":  user,
				})
				return
			}

			// Generate unique fileame
			ext := filepath.Ext(header.Filename)
			filename := uuid.New().String() + ext

			// create uploads dir if it does not exist
			os.MkdirAll("uploads", os.ModePerm)

			// save file
			out, err := os.Create("uploads/" + filename)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "create.html", gin.H{"error": "Failed to save image", "user": user})
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "create.html", gin.H{"error": "Failed to save image", "user": user})
				return
			}

			imagePath = "/uploads/" + filename
		} else {
			imagePath = ""
		}

		post := models.Post{
			Title:     title,
			Content:   content,
			ImagePath: imagePath,
			AuthorID:  user.ID,
		}

		if err := db.Create(&post).Error; err != nil {
			c.HTML(http.StatusInternalServerError, "create.html", gin.H{"error": "Failed to create post", "user": user})
			return
		}

		c.Redirect(http.StatusFound, "/posts")
	}
}

func EditPageHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetCurrentUser(c, db)
		if user == nil {
			c.Redirect(http.StatusFound, "/users/login")
			return
		}

		var post models.Post
		if err := db.First(&post, c.Param("id")).Error; err != nil {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "Post not found",
			})
			return
		}

		// Check if user owns the post
		if post.AuthorID != user.ID {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"error": "You can only edit your own posts",
			})
			return
		}

		c.HTML(http.StatusOK, "edit.html", gin.H{
			"post": post,
			"user": user,
		})
	}
}

func EditPostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetCurrentUser(c, db)
		if user == nil {
			c.Redirect(http.StatusFound, "/users/login")
			return
		}

		var post models.Post
		if err := db.First(&post, c.Param("id")).Error; err != nil {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "Post not found",
			})
			return
		}

		// Check if user owns the post
		if post.AuthorID != user.ID {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"error": "You can only edit your own posts",
			})
			return
		}

		title := c.PostForm("title")
		content := c.PostForm("content")
		removeImage := c.PostForm("removeImage") == "on"

		// Handle image removal if requested
		if removeImage && post.ImagePath != "" {
			os.Remove("." + post.ImagePath)
			post.ImagePath = ""
		}

		// Handle file upload if a new image is provided
		file, header, err := c.Request.FormFile("image")
		if err == nil {
			defer file.Close()

			if !isImage(header) {
				c.HTML(http.StatusBadRequest, "edit.html", gin.H{
					"error": "Please upload a valid image file (JPEG, PNG, GIF, or WebP)",
					"post":  post,
					"user":  user,
				})
				return
			}

			// Generate unique filename
			ext := filepath.Ext(header.Filename)
			filename := uuid.New().String() + ext

			// Create uploads directory if it doesn't exist
			os.MkdirAll("uploads", os.ModePerm)

			// Save the file
			out, err := os.Create("uploads/" + filename)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "edit.html", gin.H{
					"error": "Failed to save image",
					"post":  post,
					"user":  user,
				})
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "edit.html", gin.H{
					"error": "Failed to save image",
					"post":  post,
					"user":  user,
				})
				return
			}

			// Delete old image if it exists
			if post.ImagePath != "" {
				os.Remove("." + post.ImagePath)
			}

			post.ImagePath = "/uploads/" + filename
		}

		post.Title = title
		post.Content = content

		db.Save(&post)
		c.Redirect(http.StatusFound, "/posts/post/"+c.Param("id"))
	}
}

func DeletePostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := middleware.GetCurrentUser(c, db)
		if user == nil {
			c.Redirect(http.StatusFound, "users/login")
			return
		}

		var post models.Post
		if err := db.First(&post, c.Param("id")).Error; err != nil {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "Post not found",
			})
			return
		}

		// Check if user owns the post
		if post.AuthorID != user.ID {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"error": "You can only delete your own posts",
			})
			return
		}

		// Delete associated image if it exists
		if post.ImagePath != "" {
			os.Remove("." + post.ImagePath)
		}

		db.Delete(&post)
		c.Redirect(http.StatusFound, "/posts")
	}
}

// Add this function to validate image files
func isImage(fileHeader *multipart.FileHeader) bool {
	allowedTypes := []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
	}

	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	// Read the first 512 bytes to detect content type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buffer)

	for _, t := range allowedTypes {
		if contentType == t {
			return true
		}
	}

	return false
}
