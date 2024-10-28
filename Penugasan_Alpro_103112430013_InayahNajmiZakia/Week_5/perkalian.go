package main

import (
	"fmt"
)

func main() {
	var n int

	// Read the number of elements
	fmt.Scanln(&n)

	// Declare a slice to store the numbers
	numbers := make([]int, n)

	// Read the numbers into the slice
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	// Multiply each number by its index (starting from 1) and print the result
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", numbers[i]*(i+1))
	}
	fmt.Println()
}
