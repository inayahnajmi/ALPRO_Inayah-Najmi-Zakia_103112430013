package main

import (
	"fmt"
)

func main() {
	var totalBelanjaAwal int
	var diskonPersen int

	// Input total belanja awal
	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&totalBelanjaAwal)

	// Input persentase diskon
	fmt.Print("Masukkan besarnya diskon (dalam persen): ")
	fmt.Scan(&diskonPersen)

	// Hitung total belanja akhir setelah diskon
	diskon := (totalBelanjaAwal * diskonPersen) / 100
	totalBelanjaAkhir := totalBelanjaAwal - diskon

	// Output total belanja akhir
	fmt.Printf("Total belanja akhir setelah diskon: %d\n", totalBelanjaAkhir)
}
