package main

import (
	"fmt"
)

// Fungsi untuk menghitung penjumlahan dari 1 sampai n
func jumlahBilangan(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	hasil := jumlahBilangan(n)
	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah: %d\n", n, hasil)
}
