package main

import (
	"fmt"
)

func isPerfectNumber(angka int) bool {
	sumOfFactors := 0

	for i := 1; i <= angka/2; i++ {
		if angka%i == 0 {
			sumOfFactors += i
		}
	}
	return sumOfFactors == angka
}

func main() {
	var input int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&input)

	if isPerfectNumber(input) {
		fmt.Printf("Ya, %d adalah bilangan sempurna.\n", input)
	} else {
		fmt.Printf("Tidak, %d bukan bilangan sempurna.\n", input)
	}
}
