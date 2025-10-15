package rendertemplates

import "github.com/gin-contrib/multitemplate"

func SetupAllTemplates() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	base := "ui/html/pages/base.html"
	snippetPath := "ui/html/pages/snippets/"
	authPath := "ui/html/pages/auth/"

	renderer.AddFromFiles("home.html", base, snippetPath+"home.html")
	renderer.AddFromFiles("create.html", base, snippetPath+"create.html")
	renderer.AddFromFiles("view.html", base, snippetPath+"view.html")
	renderer.AddFromFiles("error.html", base, "ui/html/pages/error.html")

	renderer.AddFromFiles("register.html", base, authPath+"register.html")
	renderer.AddFromFiles("login.html", base, authPath+"login.html")
	renderer.AddFromFiles("profile.html", base, authPath+"profile.html")

	return renderer
}
