package main

import (
	"fmt"
)

func main() {
	var n, a, t int
	var luas float64

	fmt.Print("Masukkan jumlah segitiga: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		// Meminta input panjang alas dan tinggi untuk setiap segitiga
		fmt.Printf("Masukkan panjang alas segitiga ke-%d: ", i)
		fmt.Scan(&a)
		fmt.Printf("Masukkan tinggi segitiga ke-%d: ", i)
		fmt.Scan(&t)

		// Menghitung luas segitiga
		luas = 0.5 * float64(a) * float64(t)

		// Menampilkan hasil luas segitiga
		fmt.Printf("Luas segitiga ke-%d adalah: %.2f\n", i, luas)
	}
}