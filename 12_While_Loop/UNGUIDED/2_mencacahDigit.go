package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	if number <= 0 {
		fmt.Println("Masukan harus bilangan bulat positif.")
		return
	}

	fmt.Println("Digit dari kanan ke kiri:")
	for number > 0 {
		digit := number % 10
		fmt.Println(digit)
		number /= 10
	}
}
