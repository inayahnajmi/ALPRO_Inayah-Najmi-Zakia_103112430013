package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	maxDigit := 0
	for number > 0 {
		digit := number % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		number /= 10
	}
	fmt.Printf("Digit terbesar adalah: %d\n", maxDigit)
}
