package main

import (
	"flag"
	"log"
	"os"

	"github.com/Script-Savant/Golang/snippet-box/cmd/web/config"
	rendertemplates "github.com/Script-Savant/Golang/snippet-box/cmd/web/renderTemplates"
	"github.com/Script-Savant/Golang/snippet-box/cmd/web/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "user:password@tcp(localhost:3306)/snippetbox_db?parseTime=True&loc=Local", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	config.OpenDB(*dsn)

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	router.HTMLRender = rendertemplates.SetupAllTemplates()

	router.Static("/ui/static", "./ui/static")

	routes.SetupRoutes(router)
	routes.SetupAuthRoutes(router)

	infoLog.Printf("Starting server on %s", *addr)
	if err := router.Run(*addr); err != nil {
		errorLog.Fatal(err)
	}
}
