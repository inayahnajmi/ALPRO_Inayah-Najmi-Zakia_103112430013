package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung volume kerucut
func volumeKerucut(jariJari, tinggi float64) float64 {
	return (1.0 / 3.0) * math.Pi * math.Pow(jariJari, 2) * tinggi
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	// Memproses input dan menghitung volume untuk setiap kerucut
	for i := 0; i < n; i++ {
		var jariJari, tinggi float64
		fmt.Print("Masukkan jari-jari dan tinggi kerucut: ")
		fmt.Scan(&jariJari, &tinggi)

		volume := volumeKerucut(jariJari, tinggi)
		fmt.Printf("Volume kerucut %d adalah: %.12f\n", i+1, volume)
	}
}
