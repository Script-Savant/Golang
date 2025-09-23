package utils

import (
	

	"github.com/gin-contrib/multitemplate"
)

func SetupTemplates() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	renderer.AddFromFiles("dashboard.html", "templates/base.html", "templates/dashboard.html")
	renderer.AddFromFiles("login.html", "templates/base.html", "templates/login.html")
	renderer.AddFromFiles("register.html", "templates/base.html", "templates/register.html")
	renderer.AddFromFiles("send.html", "templates/base.html", "templates/send.html")
	renderer.AddFromFiles("success.html", "templates/base.html", "templates/success.html")

	return renderer
}