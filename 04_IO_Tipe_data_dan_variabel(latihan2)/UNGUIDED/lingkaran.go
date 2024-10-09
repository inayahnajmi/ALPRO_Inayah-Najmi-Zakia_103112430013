package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	// Input jari-jari lingkaran
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	// Menghitung luas dan keliling lingkaran
	luas := math.Pi * r * r
	keliling := 2 * math.Pi * r

	// Menampilkan hasil
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
