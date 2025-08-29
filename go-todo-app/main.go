package main

import (
	"fmt"
	"go-todo-app/config"
	"go-todo-app/models"
	"html/template"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{db}
}

func main() {
	db := config.DatabaseConnection()
	tc := NewTodoController(db)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/todos", tc.todosHandler)
	http.HandleFunc("/api/delete-todo", tc.deleteTodo)
	http.HandleFunc("/api/complete-todo", tc.completeTodo)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Unable to load index", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func renderTodoHTML(todo models.Todo) string {
	completedStatus := ""
	bgColor := "white"
	buttonText := "Complete"

	if todo.Completed {
		completedStatus = " (completed)"
		bgColor = "#f0f0f0"
		buttonText = "Uncomplete"
	}

	return fmt.Sprintf(`
		<div class="todo-item" id="todo-%d" style="background-color: %s;">
			<p><strong>%s</strong>%s</p>
			<button hx-post="/api/delete-todo"
			        hx-target="#todo-%d"
			        hx-swap="outerHTML"
			        hx-include="#todo-%d [name=id]"
			        type="button">Delete</button>
			<button hx-post="/api/complete-todo"
			        hx-target="#todo-%d"
			        hx-swap="outerHTML"
			        hx-include="#todo-%d [name=id]"
			        type="button">%s</button>
			<input type="hidden" name="id" value="%d">
		</div>`, todo.ID, bgColor, todo.Title, completedStatus,
		todo.ID, todo.ID, todo.ID, todo.ID, buttonText, todo.ID)
}

func (tc *TodoController) todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tc.getTodos(w, r)
	case http.MethodPost:
		tc.addTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (tc *TodoController) getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	if err := tc.db.Find(&todos).Error; err != nil {
		http.Error(w, "Unable to fetch todo items", http.StatusInternalServerError)
		return
	}

	var builder strings.Builder
	for _, todo := range todos {
		builder.WriteString(renderTodoHTML(todo))
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(builder.String()))
}

func (tc *TodoController) addTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	todo := models.Todo{Title: title, Completed: false}
	if err := tc.db.Create(&todo).Error; err != nil {
		http.Error(w, "Unable to add todo item", http.StatusInternalServerError)
		return
	}

	html := renderTodoHTML(todo)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func (tc *TodoController) deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := tc.db.Delete(&models.Todo{}, id).Error; err != nil {
		http.Error(w, "Unable to delete todo item", http.StatusInternalServerError)
		return
	}

	// return empty response (HTMX will swap it out)
	w.Write([]byte(""))
}

func (tc *TodoController) completeTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var todo models.Todo
	if err := tc.db.First(&todo, id).Error; err != nil {
		http.Error(w, "Unable to fetch todo item", http.StatusInternalServerError)
		return
	}

	// ✅ toggle status
	todo.Completed = !todo.Completed

	if err := tc.db.Save(&todo).Error; err != nil {
		http.Error(w, "Error updating todo item", http.StatusInternalServerError)
		return
	}

	html := renderTodoHTML(todo)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
