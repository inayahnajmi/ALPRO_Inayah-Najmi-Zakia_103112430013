package main

import (
	"fmt"
)

func main() {
	for selesai := false; !selesai; {
		var a, b float64
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&a, &b)
		fmt.Println("Sepeda motor pak Andi akan oleng:", b-a >= 9)
		selesai = (a+b > 150) || (a < 0 || b < 0)
	}
	fmt.Println("Program selesai")
}