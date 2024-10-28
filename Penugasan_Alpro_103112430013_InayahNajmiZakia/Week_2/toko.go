package main

import (
	"fmt"
)

func main() {
	// Inisialisasi harga beli untuk tiga barang
	var harga1, harga2, harga3 float64

	// Meminta input harga dari user
	fmt.Print("Masukkan harga beli untuk barang pertama, kedua, dan ketiga: ")
	fmt.Scan(&harga1, &harga2, &harga3)

	// Menghitung harga jual dengan keuntungan 5%
	hargaJual1 := harga1 * 1.05
	hargaJual2 := harga2 * 1.05
	hargaJual3 := harga3 * 1.05

	// Menampilkan hasil
	fmt.Printf("Harga jual barang pertama: %.2f\n", hargaJual1)
	fmt.Printf("Harga jual barang kedua: %.2f\n", hargaJual2)
	fmt.Printf("Harga jual barang ketiga: %.2f\n", hargaJual3)
}
