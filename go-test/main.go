package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate" // Optional: for more complex template loading
)

func main() {

	renderer := multitemplate.NewRenderer()

	renderer.AddFromFiles("index", "templates/base.html", "templates/index.html")
	renderer.AddFromFiles("about", "templates/base.html", "templates/about.html")


	router := gin.Default()

	// Option 1: Simple loading with LoadHTMLGlob (ensure base.html is also matched)
	router.HTMLRender = renderer

	// Option 2: Using gin-contrib/multitemplate for more control
	// router.HTMLRender = loadTemplates("./templates")

	router.GET("/", func(c *gin.Context) {
		fmt.Println(">>> Load index page")
		c.HTML(200, "index", gin.H{}) // Render the specific child template
	})

	router.GET("/about", func(c *gin.Context) {
		fmt.Println(">>> Load about page")
		c.HTML(200, "about", gin.H{}) // Render the specific child template
	})

	router.Run(":8080")
}

// func loadTemplates(templatesDir string) multitemplate.Renderer {
//     r := multitemplate.NewRenderer()
//     // ... logic to load layouts and includes as shown in gin-contrib/multitemplate example
//     return r
// }
