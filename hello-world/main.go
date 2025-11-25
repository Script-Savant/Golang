package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", indexHandler)

	router.Run(":8000")
}

type Person struct {
	FirstName string `xml:"firstname,attr"`
	LastName  string `xml:"lastname,attr"`
}

func indexHandler(c *gin.Context) {
	c.XML(200, Person{
		FirstName: "Jane",
		LastName:  "Smith",
	})
}
