package main

import (
	"fmt"
	"math"
)

func main() {
	// Koordinat titik A, B, dan C
	var ax, ay, bx, by, cx, cy float64

	// Input koordinat titik A
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scan(&ax, &ay)

	// Input koordinat titik B
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scan(&bx, &by)

	// Input koordinat titik C
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scan(&cx, &cy)

	// Hitung panjang sisi-sisi segitiga
	sisiAB := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2)) // Panjang AB
	sisiBC := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2)) // Panjang BC
	sisiCA := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2)) // Panjang CA

	// Menentukan sisi terpanjang
	terpanjang := sisiAB
	if sisiBC > terpanjang {
		terpanjang = sisiBC
	}
	if sisiCA > terpanjang {
		terpanjang = sisiCA
	}

	// Output panjang sisi terpanjang
	fmt.Printf("%.2f\n", terpanjang)
}
