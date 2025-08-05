package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Jane Smith")
	})

	fmt.Println("Starting server on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}