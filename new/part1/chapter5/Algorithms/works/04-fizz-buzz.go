package works

import "fmt"

func FizzBuzz(n int){
    for i := 1; i < n; i++{
        if i % 3 == 0 && i % 5 == 0{
            fmt.Print("Fizz Buzz, ")
        } else if i % 3 == 0 {
            fmt.Print("Fizz, ")
        } else if i % 5 == 0 {
            fmt.Print("Buzz, ")
        } else {
            fmt.Print(i, ", ")
        }
    }
    if n % 3 == 0 && n % 5 == 0{
        fmt.Print("Fizz Buzz")
    } else if n % 3 == 0 {
        fmt.Print("Fizz")
    } else if n % 5 == 0 {
        fmt.Print("Buzz")
    } else {
        fmt.Print(n)
    }
}
