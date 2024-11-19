package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("Masukkan bilangan bulat (b > 0): ")
	fmt.Scan(&b)

	if b <= 0 {
		fmt.Println("Bilangan harus lebih besar dari 0!")
		return
	}

	fmt.Print("Faktor: ")
	faktor := []int{}
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	isPrima := len(faktor) == 2 
	fmt.Println("Prima:", isPrima)
}
