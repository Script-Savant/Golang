package handlers

import (
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/Script-Savant/Golang/snippet-box/cmd/web/utils"
	"github.com/Script-Savant/Golang/snippet-box/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db}
}

func (ac *AuthController) Register(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", gin.H{})
		return
	}

	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))
	password2 := strings.TrimSpace(c.PostForm("password2"))

	if password != password2 {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Passwords do not match"})
		return
	}

	if len(password) < 8 {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Password must be at least 8 characters long"})
		return
	}

	if !containsUppercase(password) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Password must at least have 1 capital letter"})
		return
	}

	if !containsLowercase(password) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Password must at least have 1 small letter"})
		return
	}

	if !containsNumber(password) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Password must at least have 1 number"})
		return
	}

	if !containsSpecialCharacters(password) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Password must at least have 1 special character"})
		return
	}

	// check if username exists
	var userExists models.User
	if err := ac.DB.Where("username = ?", username).First(&userExists).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Username is already taken. Try another username"})
		return
	}

	// new user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Error encrypting password"})
		return
	}

	newUser := models.User{
		Username:     username,
		PasswordHash: string(passwordHash),
	}
	if err := ac.DB.Create(&newUser).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Error registering user"})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func (ac *AuthController) Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	}

	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))

	var user models.User
	if err := ac.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"error": "Incorrect password or username"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"error": "Incorrect password or username"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/")
}

func (ac *AuthController) Profile(c *gin.Context) {
	user := utils.GetCurrentUser(c, ac.DB)

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "profile.html", gin.H{"user": user})
	}

	oldPass := strings.TrimSpace(c.PostForm("oldPass"))
	newPass := strings.TrimSpace(c.PostForm("newPass"))
	newPass2 := strings.TrimSpace(c.PostForm("newPass2"))

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPass)); err != nil {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Incorrect password"})
		return
	}

	if newPass != newPass2 {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Passwords do not match"})
		return
	}

	if len(newPass) < 8 {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Password must be at least 8 characters long"})
		return
	}

	if !containsUppercase(newPass) {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Password must at least have 1 capital letter"})
		return
	}

	if !containsLowercase(newPass) {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Password must at least have 1 small letter"})
		return
	}

	if !containsNumber(newPass) {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Password must at least have 1 number"})
		return
	}

	if !containsSpecialCharacters(newPass) {
		c.HTML(http.StatusBadRequest, "profile.html", gin.H{"error": "Password must at least have 1 special character"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "profile.html", gin.H{"error": "Error encrypting new password"})
		return
	}

	user.PasswordHash = string(hashedPass)

	if err := ac.DB.Save(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "profile.html", gin.H{"error": "Error updating user password"})
		return
	}

	c.Redirect(http.StatusFound, "/profile")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

func containsUppercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func containsLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func containsNumber(s string) bool {
	re := regexp.MustCompile(`\d`)

	return re.MatchString(s)
}

func containsSpecialCharacters(s string) bool {
	specialCharacters := "!@#$%^&*"

	return strings.ContainsAny(s, specialCharacters)
}
