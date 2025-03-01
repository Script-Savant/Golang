package main

import (
	"fmt"
)

func exercise1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	names := []string{"Jim", "Jane", "Joe", "June"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func exercise3() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.24.0",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}

func exercise4() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  7,
		"Never": 1,
		"Up":    4,
	}

	var count int
	var word string

	for key, value := range words {
		if value > count {
			count = value
			word = key
		}
	}
	fmt.Println("Most popular word :", word)
	fmt.Println("With a count of :", count)
}

func exercise5() {
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(num)
		}
	}
}

func main() {
	exercise5()
}
