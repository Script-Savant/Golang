package handlers

import (
	"go-stripe/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h *AuthHandler) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Invalid form data"})
		return
	}

	var existingUser models.User
	if err := h.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Email already registered"})
		return
	}

	if err := user.HashPassword(); err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Error creating user"})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Error creating user"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *AuthHandler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	if err := h.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"error": "Invalid email or password"})
		return
	}

	if err := user.CheckPassword(password); err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"error": "Invalid email or password"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

func (h *AuthHandler) Dashboard(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var user models.User
	h.DB.First(&user, userID)

	var transactions []models.Transaction
	h.DB.Where("user_id = ?", userID).Order("created_at desc").Limit(10).Find(&transactions)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{"user": user, "transactions": transactions})
}
