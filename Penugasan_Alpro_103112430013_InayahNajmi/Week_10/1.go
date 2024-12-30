package main

import "fmt"

func main() {
	var a, b, c int

	// Input tiga bilangan bulat positif
	fmt.Print("Masukkan tiga sisi segitiga (dipisahkan spasi): ")
	fmt.Scan(&a, &b, &c)

	// Cek jenis segitiga
	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || b == c || a == c {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}