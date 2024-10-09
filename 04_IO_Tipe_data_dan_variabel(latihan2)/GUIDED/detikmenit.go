package main

import (
	"fmt"
)

func main() {
	// Input jumlah detik
	var detik int
	fmt.Print("Masukkan jumlah detik: ")
	fmt.Scan(&detik)

	// Menghitung jam, menit, dan detik
	jam := detik / 3600
	menit := (detik % 3600) / 60
	remainingdetik := detik % 60

	// Menampilkan hasil
	fmt.Printf("%d jam %d menit dan %d detik\n", jam, menit, remainingdetik)
}
