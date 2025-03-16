package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connnect() {
	conn, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/booksdb?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
