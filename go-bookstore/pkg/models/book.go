package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)
var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) (*Book, error) {
	result := db.Create(book)
	return book, result.Error
}

func UpdateBookByID(id int64, updatedBook *Book) (*Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	book.Name = updatedBook.Name
	book.Author = updatedBook.Author
	book.Publication = updatedBook.Publication
	db.Save(&book)
	return &book, nil
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return  Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}