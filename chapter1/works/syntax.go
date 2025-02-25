package works

import (
	"fmt"
	"reflect"
)

func CheckTypes() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello Go"))
	fmt.Println(reflect.TypeOf('A'))
}

func VariableDeclarations() {
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jersk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left")
}

func Conversion(){
	length := 1.2
	width := 2
	fmt.Println("Area is", length * float64(width))
	fmt.Println("Length > width?", length > float64(width))
}

func Exercise(){
	price := 100
	fmt.Println("Price is", price, "dollars.")
	
	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	
	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")
	
	availableFunds := 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}