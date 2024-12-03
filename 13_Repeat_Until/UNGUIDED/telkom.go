package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scanln(&input)

		if input == "telkom" || input == "Telkom" || input == "TELKOM" || input == "TeLkOm" {
			fmt.Println("Program selesai.")
			break
		} else {
			fmt.Println("Anda mengetik:", input)
		}
	}
}
