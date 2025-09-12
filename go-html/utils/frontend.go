package utils

import (
	"text/template"
	"time"

	"github.com/gin-contrib/multitemplate"
)

func formatDate(t time.Time) string {
	return t.Format("Monday 2, 2006")
}

func LoadTemplates() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	parseTime := template.FuncMap{
		"formatDate": formatDate,
	}

	// Posts pages (use base.html)
	renderer.AddFromFilesFuncs("index.html", parseTime, "templates/base.html", "templates/index.html")
	renderer.AddFromFilesFuncs("post.html", parseTime, "templates/base.html", "templates/post.html")
	renderer.AddFromFiles("create.html", "templates/base.html", "templates/create.html")
	renderer.AddFromFiles("edit.html", "templates/base.html", "templates/edit.html")
	renderer.AddFromFiles("error.html", "templates/base.html", "templates/error.html")

	// Auth pages (use auth.html)
	renderer.AddFromFiles("login.html", "templates/base.html", "templates/login.html")
	renderer.AddFromFiles("register.html", "templates/base.html", "templates/register.html")

	return renderer
}
