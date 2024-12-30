package main

import "fmt"

func main() {
	const correctUsername = "inay00"
	const correctPassword = "inaaay"
	var failedAttempts int
	var username, password string

	fmt.Println("Masukkan username dan password:")

	for {
		fmt.Print("Username: ")
		fmt.Scanln(&username)

		fmt.Print("Password: ")
		fmt.Scanln(&password)

		if username == correctUsername && password == correctPassword {
			fmt.Println("Login berhasil.")
			break
		} else {
			fmt.Println("Login gagal. Coba lagi.")
			failedAttempts++
		}
	}

	fmt.Printf("Jumlah login gagal: %d\n", failedAttempts)
}
