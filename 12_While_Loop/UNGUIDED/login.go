package main

import (
	"fmt"
)

func main() {
	const passwordBenar = "inay123"

	const percobaan = 3

	var password string
	attempts := 0

	fmt.Println("Login")

	for attempts < percobaan {
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if password == passwordBenar {
			fmt.Println("Login berhasil!")
			return
		} else {
			attempts++
			fmt.Printf("Password salah! Anda memiliki %d kesempatan lagi.\n", percobaan-attempts)
		}
	}

	fmt.Println("Login ditolak, Anda terlalu banyak mencoba")
}
