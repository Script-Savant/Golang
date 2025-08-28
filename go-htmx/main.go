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
	<script src="https://unpkg.com/htmx.org"></script>
	</head>
	<body>
	<h1>HTMX demo</h1>
	<div id="content">
		<p>Click the button to fetch updated content!</p>
	</div>
	<button hx-get="/update" hx-target="#content" hx-swap="innerHTML">
		Get updated content
	</button>	
	</body>
	</html>
	`)
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<p>Content updated at: `+r.RemoteAddr+`</p>`)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/update", update)

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
