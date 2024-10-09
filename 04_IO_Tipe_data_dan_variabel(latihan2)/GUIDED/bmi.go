package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk berat badan dan tinggi badan
	var berat, tinggi, bmi float64

	// Input berat badan (dalam kg)
	fmt.Print("Masukkan berat badan: ")
	fmt.Scan(&berat)

	// Input tinggi badan (dalam meter)
	fmt.Print("Masukkan tinggi badan: ")
	fmt.Scan(&tinggi)

	// Menghitung BMI
	bmi = berat / (tinggi * tinggi)

	// Menampilkan hasil BMI
	fmt.Printf("Nilai BMI Anda adalah: %.2f\n", bmi)
}
