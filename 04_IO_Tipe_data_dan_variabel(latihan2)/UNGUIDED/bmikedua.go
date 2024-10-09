package main

import (
	"fmt"
)

func main() {
	var bmi float64
	var tinggiBadan float64

	// Input nilai BMI
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)

	// Input tinggi badan dalam meter
	fmt.Print("Masukkan tinggi badan (dalam meter): ")
	fmt.Scan(&tinggiBadan)

	// Hitung berat badan
	beratBadan := bmi * (tinggiBadan * tinggiBadan)

	// Output berat badan
	fmt.Printf("Berat badan seseorang adalah: %.2f kg\n", beratBadan)
}