package controllers

/*
1. initiate a database instance
2. create posts
3. get all posts
4. get a single post by id
5. update a post
7. delete a post
8. Like a post
9. share a post
*/

import (
	"go-blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PostController -> handle post related ops
type PostController struct {
	DB *gorm.DB
}

// create a PostController with database connection
func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

// create a new blog post
func (pc *PostController) CreatePost(c *gin.Context) {
	/*
		- get the user email from context
		- find the user by email to get their id
		- bind the incoming json to a Post struct
		- set the author id to the current user id
		- create the post in the db
		- return success response
	*/

	// 1. get the user email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	// 2. find the user by email to get their id
	var user models.User
	if err := pc.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 3. bind the incoming json to a post struct
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. set the author id to the current user id
	post.AuthorID = user.ID

	// 5. create the post in the database
	if err := pc.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
		return
	}

	// return success
	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}
