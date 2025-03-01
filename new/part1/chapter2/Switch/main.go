package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Monday

	switch dayBorn{
	case time.Saturday, time.Sunday:
		fmt.Println("Born on a weekend")
	default:
		fmt.Println("Born on a weekday")
	}
}