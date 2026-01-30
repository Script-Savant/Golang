package handlers

import (
	"net/http"
	"url-shortener/internal/config"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ShortenURL(c *gin.Context) {
	original := c.PostForm("url")

	shortcode := uuid.New().String()[:6]

	url := models.URL{
		OriginalURL: original,
		ShortCode:   shortcode,
	}

	config.DB.Create(&url)

	c.HTML(http.StatusOK, "result", gin.H{"ShortURL": c.Request.Host + "/" + shortcode})
}
