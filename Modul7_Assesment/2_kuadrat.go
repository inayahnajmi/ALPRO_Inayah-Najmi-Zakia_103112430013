package main

import (
	"fmt")

func main(){
	var N int

	fmt.Print("masukan nilai n : ")
	fmt.Scanln(&N)

	if N <= 0 {
		fmt.Print("nilai n harus bilangan bulat.")
		return
	}

	fmt.Println("kuadrat dari bilangan 1 sampai", N, ":")
	for i := 1; i <= N; i++ {
		square := i * i
		fmt.Printf("kuadrat dari %d adalah: %d\n", i, square)
	}


}