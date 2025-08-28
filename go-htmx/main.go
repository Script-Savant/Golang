package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintln(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Hello, World</title>
	</head>
	<body>
	<h1>Hello, World!</h1>
	<p>Welcome to your first Go web server.</p>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
