package main

import (
    "fmt"
)

func main() {
    var fahrenheit float64

    // Input suhu dalam Fahrenheit
    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&fahrenheit)

    // Konversi Fahrenheit ke Kelvin
    kelvin := (fahrenheit - 32) * 5 / 9 + 273.15

    // Menampilkan hasil
    fmt.Printf("Suhu dalam Kelvin adalah: %.2f\n", kelvin)
}
