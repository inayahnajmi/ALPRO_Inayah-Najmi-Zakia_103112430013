package main

import (
	"fmt"
)

func main() {
	var numbers [10]int
	fmt.Println("Masukkan 10 angka (angka pertama sebagai x, sisanya 0 atau x):")
	for i := 0; i < 10; i++ {
		fmt.Scan(&numbers[i])
	}

	x := numbers[0]

	countX := 0
	count0 := 0
	for i := 1; i < 10; i++ {
		if numbers[i] == x {
			countX++
		} else if numbers[i] == 0 {
			count0++
		}
	}

	if countX > count0 {
		fmt.Printf("Modus = %d\n", x)
	} else {
		fmt.Println("Modus = 0")
	}
}