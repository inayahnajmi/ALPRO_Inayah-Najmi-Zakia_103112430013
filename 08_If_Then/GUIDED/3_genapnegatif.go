package main

import (
	"fmt"
)

func main() {
	var number int

	// Meminta input bilangan dari pengguna
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&number)

	// Memeriksa apakah bilangan genap negatif dengan if-else
	var isEvenNegative bool
	if number < 0 && number%2 == 0 {
		isEvenNegative = true
	} else {
		isEvenNegative = false
	}

	// Menampilkan hasil
	fmt.Println(isEvenNegative)
}
