package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Masukkan satu karakter: ")
	fmt.Scan(&input)

	if len(input) == 0 {
		fmt.Println("Input tidak valid")
		return
	}

	karakter := input[0]

	if (karakter >= 'A' && karakter <= 'Z') || (karakter >= 'a' && karakter <= 'z') {
		switch karakter {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			fmt.Println("bukan konsonan")
		default:
			fmt.Println("konsonan")
		}
	} else {
		fmt.Println("bukan konsonan")
	}
}
