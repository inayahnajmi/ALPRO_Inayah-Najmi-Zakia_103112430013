package main

import "fmt"

func main() {
	var input float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	// Pembulatan ke atas secara manual
	ceil := float64(int(input))
	if input > ceil {
		ceil += 1
	}

	current := input
	fmt.Println("Hasil perulangan hingga mencapai pembulatan ke atas:")

	for current < ceil {
		current += 0.1
		fmt.Printf("%.1f\n", current)
	}
}
