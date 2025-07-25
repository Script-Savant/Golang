package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	allBooks := models.GetAllBooks()
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	book := &models.Book{}
	utils.ParseBody(r, book)

	newBook, err := models.CreateBook(book)
	if err != nil {
		http.Error(w, "Could not create book", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	updatedData := &models.Book{}
	utils.ParseBody(r, updatedData)

	updatedBook, err := models.UpdateBookByID(id, updatedData)
	if err != nil {
		http.Error(w, "Could not update book", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}