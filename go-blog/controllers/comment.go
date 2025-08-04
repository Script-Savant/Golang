package controllers

/*
Handle comment related operations
- a database connection instance
- create comments
- get all comments
- get comment
- like comment
*/

import (
	"go-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// handle comment related ops
type CommentController struct {
	DB *gorm.DB
}

// create a new comment controller with db connection
func NewCommentcontroller(db *gorm.DB) *CommentController {
	return &CommentController{DB: db}
}

// create anew comment
func (cc *CommentController) CreateComment(c *gin.Context) {
	/*
		- get the user email from context
		- find the user by email to get their id
		- parse the post id from the url
		- bind incoming json for the comment
		- set the post and user id for the comment
		- create the comment in the db
		- return success and created comment
	*/

	// 1. get the user email from context
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	// 2. Find the user by email to get their id
	var user models.User
	if err := cc.DB.Where("email = ?", email.(string)).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 3. parse the post id from the url
	str := c.Param("postId")
	postID, err := strconv.ParseUint(str, 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing post id"})
		return
	}

	// 4. bind the incoming json to a comment struct
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 5. attach post id and user id to the comment
	comment.PostID = uint(postID)
	comment.UserID = user.ID

	// 6. create the comment in the db
	if err := cc.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating comment"})
		return
	}

	// 7. return success with comment created
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment created successfully",
		"comment": comment,
	})
}

// reterieve all comments for a post with pagination
func (cc *CommentController) GetComments(c *gin.Context) {
	/*
		- parse the post id from url
		- parse pgaination parameters
		- query comments & user info
		- return the comments
	*/

	// 1. parse the post id from the url
	postID := c.Param("postId")

	// 2. parse pagination params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// 3. Query comments with pagination,including user information
	var comments []models.Comment
	if err := cc.DB.Preload("User").Where("post_id = ?", postID).Offset(offset).Limit(limit).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching comments"})
		return
	}

	// 4. return the comments
	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"page":     page,
		"limit":    limit,
	})
}
