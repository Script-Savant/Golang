package main

import (
	"fmt"
	"mongo-golang/config"
	"mongo-golang/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/users", routes.CreateUser).Methods("POST")
	router.HandleFunc("/users", routes.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
