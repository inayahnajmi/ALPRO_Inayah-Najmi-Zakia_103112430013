package main

import (
	"fmt"
)

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	if isPositive(angka) {
		fmt.Printf("Bilangan %d adalah positif.\n", angka)
	}else {
		fmt.Printf("Bilangan %d bukan positif.\n", angka)
	}
}

func isPositive(n int) bool {
	return n > 0
}
