package main

import (
	"fmt"
)

func main() {
	var a, b int

	// Input dua bilangan bulat a dan b
	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan b: ")
	fmt.Scan(&b)

	// Validasi: pastikan a <= b
	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	// Tampilkan baris bilangan dari a sampai b
	fmt.Printf("Baris bilangan dari %d sampai %d: ", a, b)
	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println() // Baris baru di akhir output
}
