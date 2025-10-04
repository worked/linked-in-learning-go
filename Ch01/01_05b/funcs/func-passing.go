package main

import "fmt"

// ProcessItems processes a slice of integers using a given operation
func ProcessItems(items []int, operation func(int) int) []int {
	result := make([]int, len(items))
	for i, item := range items {
		result[i] = operation(item)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Double each number
	doubledNumbers := ProcessItems(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled numbers:", doubledNumbers) // Output: Doubled numbers: [2 4 6 8 10]

	// Square each number
	squaredNumbers := ProcessItems(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squared numbers:", squaredNumbers) // Output: Squared numbers: [1 4 9 16 25]
}
