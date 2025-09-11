package controllers

import (
	"go-html/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginPageHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	// If user is already logged in, redirect to home
	if userID != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"error": c.Query("error"),
	})
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		var user models.User
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid credentials"})
			return
		}

		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Save()

		c.Redirect(http.StatusFound, "/")
	}
}

func RegisterPageHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	// If user is already logged in, redirect to home
	if userID != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "register.html", gin.H{
		"error": c.Query("error"),
	})
}

func RegisterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")

		var user models.User
		if err := db.Where("username = ?", username).First(&user).Error; err == nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Username already taken"})
			return
		}

		if err := db.Where("email = ?", email).First(&user).Error; err == nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "A user with that email is already registered"})
			return
		}

		/*
			Ensure password security
		*/

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "An internal error occured"})
			return
		}

		newUser := models.User{
			Username: username,
			Email:    email,
			Password: string(hashedPassword),
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "Error creating new user"})
			return
		}

		c.HTML(http.StatusOK, "register.html", gin.H{"message": "User created successfully"})
		c.Redirect(http.StatusFound, "/login")
	}
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
