package routes

import (
	"url-shortener/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", nil)
	})

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.Redirect)
}
