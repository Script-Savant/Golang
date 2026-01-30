package handlers

import (
	"url-shortener/internal/config"
	"url-shortener/internal/models"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	code := c.Param("code")

	var url models.URL
	result := config.DB.Where("short_code = ?", code).First(&url)
	if result.Error != nil {
		c.AbortWithStatus(404)
		return
	}

	config.DB.Model(&url).Update("clicks", url.Clicks+1)

	c.Redirect(302, url.OriginalURL)
}
