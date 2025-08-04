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
