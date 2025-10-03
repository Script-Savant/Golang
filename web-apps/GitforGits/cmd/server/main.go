package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Config map[string]string
}

type Configuration struct {
	ServerAddress string
	DbUser string
	DbPassword string
	DbHost string
	DbName string
}

func loadConfiguration() Configuration {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, assuming production environment with OS level environment variables")
	}
	return  Configuration{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DbUser: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost: os.Getenv("DB_HOST"),
		DbName: os.Getenv("DB_NAME"),
	}
}






func (a *App) Initialize(config map[string]string) {
	// database connection logic
	connectionString := config["database"]
	var err error

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	a.DB = db

	// initialize router and routes
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {

}

func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
}

func main() {

	// config := loadConfiguration()

	config := map[string]string{
		"database": "alex:password@/gitforgits_db",
	}

	app := &App{}
	app.Initialize(config)
	app.Run(":8080")
}
