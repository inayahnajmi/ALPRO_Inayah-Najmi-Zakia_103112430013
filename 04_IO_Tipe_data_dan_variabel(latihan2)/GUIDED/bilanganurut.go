package main

import (
	"fmt"
)

func main() {
	// Input bilangan tiga digit
	var angka int
	fmt.Print("Masukkan bilangan tiga digit: ")
	fmt.Scan(&angka)

	// Memastikan input antara 100 sampai 999
	if angka < 100 || angka > 999 {
		fmt.Println("Bilangan harus antara 100 hingga 999")
		return
	}

	// Memisahkan digit-digit bilangan
	ratusan := angka / 100       // Digit ratusan
	puluhan := (angka / 10) % 10     // Digit puluhan
	satuan := angka % 10            // Digit satuan

	// Memeriksa apakah digit terurut membesar
	if ratusan < puluhan && puluhan < satuan {
		fmt.Println(true) // Bilangan terurut membesar
	} else {
		fmt.Println(false) // Bilangan tidak terurut membesar
	}
}
