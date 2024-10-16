package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Input dua bilangan bulat positif
	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&b)

	// Validasi: pastikan kedua bilangan lebih besar dari 0
	if a <= 0 || b <= 0 {
		fmt.Println("Kedua bilangan harus positif.")
		return
	}

	// Hitung perkalian dengan penjumlahan berulang
	result := 0
	for i := 0; i < b; i++ {
		result += a
	}

	// Tampilkan hasil perkalian
	fmt.Printf("Hasil perkalian %d dan %d adalah: %d\n", a, b, result)
}
