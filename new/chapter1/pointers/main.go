package main

import (
	"fmt"
	"time"
)

func exercise1() {
	var count1 *int

	count2 := new(int)

	countTemp := 5

	count3 := &countTemp

	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t)
}

func exercise2() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Timer{}

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time : %#v\n", *t)
		// fmt.Printf("time : %#v\n", t.String())
	}
}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point: ", *count)
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var count int

	add5Value(count)
	fmt.Println("add5Value post:", count)

	add5Point(&count)
	fmt.Println("add5Point post:", count)

	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}
