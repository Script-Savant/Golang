package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-1-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before :", arr)
	bubbleSort(arr)
	fmt.Println("After :", arr)
}
