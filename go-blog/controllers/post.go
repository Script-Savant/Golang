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

// GetPost retrieves a single post by ID
func (pc *PostController) GetPost(c *gin.Context) {
	// Step 1: Parse the post ID from the URL
	postID := c.Param("id")

	// Step 2: Query the post with related data (author, tags, comments)
	var post models.Post
	if err := pc.DB.Preload("Author").Preload("Tags").Preload("Comments").Preload("Comments.User").First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Step 3: Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
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

// Delete a post
func (pc *PostController) DeletePost(c *gin.Context) {
	/*
		- get the user's email from context
		- parse the post id from url
		- find the post to delete
		- verify the current user is the author of the post
		- delete the post
		- return success response
	*/

	// 1. get the user email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	// 2. parse the post id from url
	postID := c.Param("id")

	// 3. find the post to delete
	var post models.Post
	if err := pc.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 4. verify the current user is the post author
	var user models.User
	if err := pc.DB.Where("email = ?", email.(string)).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if post.AuthorID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own post"})
		return
	}

	// 5. delete the post
	if err := pc.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
		return
	}

	// 6. return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

// Liking or disliking a post
func (pc *PostController) LikePost(c *gin.Context) {
	/*
		- get the user email from the contex
		- parse the post id from the url
		- find the user
		- check if the post exists
		- determine if this is a like or dislike
		- check if the user has already liked/disliked the post
		- return successresponse
	*/

	// 1. get the user email
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	// 2. parse the post id and action from url
	postId := c.Param("id")
	action := c.Param("action")

	// 3. find the user
	var user models.User
	if err := pc.DB.Where("email = ?", email.(string)).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 4. check if the post exists
	var post models.Post
	if err := pc.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 5. determine like or dislike
	isLike := true
	if action == "dislike" {
		isLike = false
	}

	// 6. check if user already liked or disliked this post
	var existingLike models.Like
	err := pc.DB.Where("user_id = ? AND post_id = ?", user.ID, post.ID).First(&existingLike).Error
	if err == nil {
		// like exists, update it
		if existingLike.IsLike == isLike {
			// same action, remove the like
			if err := pc.DB.Delete(&existingLike).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating like"})
				return
			}
		} else {
			// different action, update the like
			existingLike.IsLike = isLike
			if err := pc.DB.Save(&existingLike).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating like"})
				return
			}
		}
	} else {
		// no existing like, create new one
		newLike := models.Like{
			UserID: user.ID,
			PostID: &post.ID,
			IsLike: isLike,
		}
		if err := pc.DB.Create(&newLike).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating like"})
			return
		}
	}

	// 7. return success response
	c.JSON(http.StatusOK, gin.H{"message": "Post reaction updated successfully"})
}

// share post
func (pc *PostController) SharePost(c *gin.Context) {
	/*
		- parse the post id from the url
		- find the post
		- increment the share count
		- return success response
	*/

	// 1. parse the post id from the url
	postID := c.Param("id")

	// 2. find the post
	var post models.Post
	if err := pc.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 3. increment the share count
	if err := pc.DB.Model(&post).Update("shares", post.Shares+1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating share count"})
		return
	}

	// 4. return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Post share count updated",
		"shares":  post.Shares + 1,
	})
}
