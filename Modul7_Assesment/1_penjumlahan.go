package main

import (
	"fmt"
)

func main() {
	var x, y int

	// input dari user
	fmt.Print("masukkan nilai x : ")
	fmt.Scanln(&x)
	fmt.Print("masukkan nilai y (x <= y): ")
	fmt.Scanln(&y)

	// x <= y
	if x > y {
		fmt.Println("nilai x harus kurang dari atau sama dengan y")
		return
	}

	// hitung jumlah dari x sampai y
	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}

	// output
	fmt.Printf("jumlah bilangan dari %d sampai %d adalah: %d\n", x, y, sum)
}