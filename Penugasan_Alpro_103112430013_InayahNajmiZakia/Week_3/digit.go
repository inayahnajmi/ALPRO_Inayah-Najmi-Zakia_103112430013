package main

import (
	"fmt"
)

func main() {
	// Inisialisasi variabel untuk menyimpan bilangan
	var num int

	// Meminta input bilangan dari pengguna
	fmt.Print("Masukkan bilangan tiga digit: ")
	fmt.Scan(&num)

	// Memisahkan digit-digit bilangan
	digit1 := num / 100           // Digit pertama (ratusan)
	digit2 := (num / 10) % 10      // Digit kedua (puluhan)
	digit3 := num % 10             // Digit ketiga (satuan)

	// Memeriksa apakah digit terurut mengecil
	if digit1 > digit2 && digit2 > digit3 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
