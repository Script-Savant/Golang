package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Get all movies
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Delete a movie
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Get a single movie
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// create a movie
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// update a movie
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	route := mux.NewRouter()

	// movies seed data
	movies = append(movies, Movie{ID: "1", Isbn: "terwy23", Title: "The Last Airbender", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "asdf123", Title: "Inception", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "3", Isbn: "qwer456", Title: "The Dark Knight", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "4", Isbn: "zxcv789", Title: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "5", Isbn: "poiu098", Title: "The Matrix", Director: &Director{FirstName: "Lana", LastName: "Wachowski"}})
	movies = append(movies, Movie{ID: "6", Isbn: "lkjh765", Title: "Fight Club", Director: &Director{FirstName: "David", LastName: "Fincher"}})
	movies = append(movies, Movie{ID: "7", Isbn: "mnbv432", Title: "Pulp Fiction", Director: &Director{FirstName: "Quentin", LastName: "Tarantino"}})
	movies = append(movies, Movie{ID: "8", Isbn: "ytre321", Title: "Forrest Gump", Director: &Director{FirstName: "Robert", LastName: "Zemeckis"}})
	movies = append(movies, Movie{ID: "9", Isbn: "rewq654", Title: "The Godfather", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "10", Isbn: "dfgh987", Title: "The Shawshank Redemption", Director: &Director{FirstName: "Frank", LastName: "Darabont"}})

	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", route))
}
