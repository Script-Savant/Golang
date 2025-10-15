package routes

import (
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/config"
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/handlers"
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	db := config.GetDB()
	snippets := handlers.NewSnippetController(db)

	router := r.Group("/")
	router.Use(middleware.AuthRequired())
	{
		router.GET("/", snippets.Home)
		router.GET("/snippet/create", snippets.SnippetCreate)
		router.POST("/snippet/create", snippets.SnippetCreate)
		router.GET("/snippet/view/:id", snippets.SnippetView)
		router.POST("/snippet/view/:id", snippets.SnippetView)
		router.POST("/snippet/delete/:id", snippets.SnippetDelete)
	}
}
