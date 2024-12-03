package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan kedua (y): ")
	fmt.Scan(&y)

	if x < y || y <= 0 {
		fmt.Println("Input tidak valid. Pastikan x >= y dan y > 0.")
		return
	}

	hasil := 0
	for x >= y {
		x -= y
		hasil++
	}

	fmt.Printf("Hasil integer division: %d\n", hasil)
}
