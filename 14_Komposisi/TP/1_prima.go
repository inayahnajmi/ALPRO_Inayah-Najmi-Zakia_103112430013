package main

import (
	"fmt"
)

func isPrime(angka int) bool {
	if angka < 2 {
		return false
	}
	for i := 2; i*i <= angka; i++ {
		if angka%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scanln(&n)

	fmt.Println("Bilangan prima dari 1 hingga", n, "adalah:")

	for i := 1; i <= n; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
