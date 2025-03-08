package works

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		PrintFizzBuzz(i)
		fmt.Print(", ")
	}
	PrintFizzBuzz(n)
	fmt.Println()
}

func PrintFizzBuzz(n int) {
	switch {
	case n%15 == 0:
		fmt.Print("Fizz Buzz")
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}

func FizzBuzz2(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz Buzz")
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		default:
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
