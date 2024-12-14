package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Masukan harus bilangan bulat positif!")
		return
	}

	count := 0

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			count++
		}
	}

	fmt.Printf("Terdapat %d bilangan ganjil dari 1 hingga %d.\n", count, n)
}
