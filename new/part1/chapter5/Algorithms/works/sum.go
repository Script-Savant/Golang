package works

func Sum(numbers []int) int {
    if len(numbers) == 0{
        return 0
    }
    total := 0
    for _, num := range numbers{
        total += num
    }

    return total
}

// using recursion
func Sum2(numbers []int) int{
    if len(numbers) == 0{
        return 0
    }

    return numbers[0] + Sum2(numbers[1:])
}
