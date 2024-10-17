package main

import "fmt"

// Fungsi untuk menghitung pemangkatan
func pangkat(basis, eksponen int) int {
	hasil := 1
	for i := 0; i < eksponen; i++ {
		hasil *= basis
	}
	return hasil
}

func main() {
	var basis, eksponen int
	fmt.Print("Masukkan bilangan basis: ")
	fmt.Scan(&basis)
	fmt.Print("Masukkan bilangan eksponen: ")
	fmt.Scan(&eksponen)

	// Menghitung hasil pangkat
	hasil := pangkat(basis, eksponen)
	fmt.Printf("Hasil dari %d pangkat %d adalah: %d\n", basis, eksponen, hasil)
}