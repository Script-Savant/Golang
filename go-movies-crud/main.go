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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// set content type to json then encode movies to json 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// set content type to json
	// fetch parameters from request
	// loop through the movies to find one whose id matches the one parameters and delete it
	// encode the new list to json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	/*Set the content to json
	fetch parameters from the request
	loop through movies to find the item with the id same as one provided by parameters
	encode the item to json
	*/
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	/*
	set content  to json
	create a variable movie
	Decode the request body and assign it to movie
	add the movie to movies list
	encode the movie to json
	*/
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set json content
	w.Header().Set("Content-Type", "application/json")
	// get params
	params := mux.Vars(r)
	// loop over the movies
	for index, item := range movies {
		if item.Id == params["id"] {
			// delete the movie with the id provided
			movies = append(movies[:index], movies[index+1:]...)
			// add a new movie - the movie provided
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	/*
	create a router from mux
	Define routes for get all, get by id, create, update and delete and their corresponding functions with a method for each
	wire each to router through HandleFunc
	Start server
	*/
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "234", Title: "Movie 1", Director: &Director{Firstname: "Jane", LastName: "Smith"}})
	movies = append(movies, Movie{Id: "2", Isbn: "345", Title: "Movie 2", Director: &Director{Firstname: "John", LastName: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
