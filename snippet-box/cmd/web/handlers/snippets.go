package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Script-Savant/Golang/snippet-box/cmd/web/utils"
	"github.com/Script-Savant/Golang/snippet-box/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SnippetController struct {
	DB *gorm.DB
}

func NewSnippetController(db *gorm.DB) *SnippetController {
	return &SnippetController{db}
}

func (sc *SnippetController) Home(c *gin.Context) {
	user := utils.GetCurrentUser(c, sc.DB)

	var snippets []models.Snippet

	if err := sc.DB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&snippets).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "home.html", gin.H{"error": "Failed to load any snippets."})
		return
	}

	for _, snippet := range snippets {
		if snippet.ExpiresAt.Before(time.Now()) {
			sc.DB.Delete(&snippet)
		}
	}

	data := gin.H{
		"snippets": snippets,
		"user":     user,
	}

	c.HTML(http.StatusOK, "home.html", data)
}

func (sc *SnippetController) SnippetCreate(c *gin.Context) {
	user := utils.GetCurrentUser(c, sc.DB)

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "create.html", gin.H{"user": user})
		return
	}

	customErr := ""
	data := gin.H{
		"user": user,
		"error": customErr,
	}

	// Handle POST
	title := strings.TrimSpace(c.PostForm("title"))
	content := strings.TrimSpace(c.PostForm("content"))
	expiresStr := c.PostForm("expires")

	if title == "" || content == "" {
		customErr = "Title and content are required"
		c.HTML(http.StatusBadRequest, "create.html", data)
		return
	}

	expires, err := strconv.Atoi(expiresStr)
	if err != nil || expires <= 0 {
		customErr = "Invalid expiration value"
		c.HTML(http.StatusBadRequest, "create.html", data)
		return
	}

	snippet := models.Snippet{
		Title:     title,
		Content:   content,
		ExpiresIn: expires,
		UserID:    user.ID,
	}

	if err := sc.DB.Create(&snippet).Error; err != nil {
		customErr = "Failed to create snippet"
		c.HTML(http.StatusInternalServerError, "create.html", data)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/snippet/view/%d", snippet.ID))
}

func (sc *SnippetController) SnippetView(c *gin.Context) {
	user := utils.GetCurrentUser(c, sc.DB)

	customErr := ""
	data := gin.H{
		"user": user,
		"error": customErr,
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		customErr = "Invalid snippet id"
		c.HTML(http.StatusBadRequest, "error.html", data)
		return
	}

	// fetch the snippet
	var snippet models.Snippet
	if err := sc.DB.First(&snippet, id).Error; err != nil {
		customErr = "Error locating that snippet"
		c.HTML(http.StatusNotFound, "error.html", data)
		return
	}


	if snippet.UserID != user.ID {
		customErr = "Unauthorized Access"
		c.HTML(http.StatusForbidden, "error.html", data)
		return
	}

	// view the snippet
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "view.html", gin.H{"snippet": snippet, "user": user})
		return
	}

	// update the snippet
	title := c.PostForm("title")
	content := c.PostForm("content")
	expiresStr := c.PostForm("expires")

	expires, err := strconv.Atoi(expiresStr)
	if err != nil {
		customErr = "Inappropriate value for number of days"
		c.HTML(http.StatusBadRequest, "view.html", data)
		return
	}

	snippet.Title = strings.TrimSpace(title)
	snippet.Content = strings.TrimSpace(content)
	snippet.ExpiresIn = expires

	if err := sc.DB.Save(&snippet).Error; err != nil {
		customErr = "Error updating the snippet"
		c.HTML(http.StatusInternalServerError, "view.html", data)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/snippet/view/%d", id))
}

func (sc *SnippetController) SnippetDelete(c *gin.Context) {
	user := utils.GetCurrentUser(c, sc.DB)
	snippetID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid snippet ID"})
		return
	}

	var snippet models.Snippet
	if err := sc.DB.First(&snippet, snippetID).Error; err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": "Error locating that snippet"})
		return
	}

	if snippet.UserID != user.ID {
		c.HTML(http.StatusForbidden, "error.html", gin.H{"error": "Unauthorized Access"})
		return
	}

	if err := sc.DB.Delete(&snippet).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Error deleting that snippet"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
