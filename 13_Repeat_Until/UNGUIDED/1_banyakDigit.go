package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&input)

	if input < 0 {
		fmt.Println("Masukan harus bilangan bulat positif.")
		return
	}

	digitCount := countDigits(input)
	fmt.Printf("Jumlah digit dari bilangan %d adalah %d\n", input, digitCount)
}

func countDigits(num int) int {
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}
