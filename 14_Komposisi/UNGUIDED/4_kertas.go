package main

import "fmt"

func main() {
	var bunga, pita string

	i := 0
	for selesai := false; !selesai; {
		i += 1
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)
		if bunga != "SELESAI" {
			pita += bunga + "-"
		}
		selesai = bunga == "SELESAI"
	}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", i-1)
}