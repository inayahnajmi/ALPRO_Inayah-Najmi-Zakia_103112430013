package main

import "fmt"

func main() {
	var total, price float64
	var itemName string
	var option int

	fmt.Println("Aplikasi Kasir Sederhana")

	for option != 2 {
		fmt.Println("\n1. Tambah barang")
		fmt.Println("2. Selesaikan transaksi")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&option)

		if option == 1 {
			fmt.Print("Nama barang: ")
			fmt.Scan(&itemName)
			fmt.Print("Harga barang: ")
			fmt.Scan(&price)
			total += price
			fmt.Printf("'%s' dengan harga %.2f ditambahkan.\n", itemName, price)
		} else if option != 2 {
			fmt.Println("Pilihan tidak valid.")
		}
	}

	fmt.Printf("\nTotal belanja: %.2f\n", total)
	fmt.Println("Terima kasih!")
}
