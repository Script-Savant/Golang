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
	"strconv"

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

// get all posts
func (pc *PostController) GetPosts(c *gin.Context) {
	/*
		- parse pagination parameters -> 10 items per page
		- query posts with pagination(author, tags)
		- return the posts
	*/

	// 1. pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "2"))
	offset := (page - 1) * limit

	// 2. query posts with pagination
	var posts []models.Post
	if err := pc.DB.Preload("Author").Preload("Tags").Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}

	// 3. return the posts
	c.JSON(http.StatusOK, gin.H{
		"Posts": posts,
		"page":  page,
		"limit": limit,
	})
}

// update an existing post
func (pc *PostController) UpdatePost(c *gin.Context) {
	/*
		- get the user email from context
		- parse the post id from the url
		- find the post to update
		- verify the current user is the author of the post
		- bind the update data
		- update the post
		- return success response
	*/

	// 1. get the user email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	// 2. parse the postID from url
	postID := c.Param("id")

	// 3. Find the post to update
	var post models.Post
	if err := pc.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 4. verify the current user is the author of the post
	var user models.User
	if err := pc.DB.Where("email = ?", email.(string)).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if post.AuthorID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own posts"})
		return
	}

	// 5. Bind the update data
	var updateData models.Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 6. Update the post
	if err := pc.DB.Model(&post).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updateing post"})
		return
	}

	// 7. return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}
