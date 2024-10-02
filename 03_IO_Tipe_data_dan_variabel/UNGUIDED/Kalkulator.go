package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    var operator string

    // Input angka pertama
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&num1)

    // Input operator (+, -, *, /)
    fmt.Print("Masukkan operator (penjumlahan, pengurangan, perkalian, pembagian): ")
    fmt.Scanln(&operator)

    // Input angka kedua
    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&num2)

    // Proses perhitungan
    switch operator {
    case "penjumlahan":
        fmt.Printf("Hasil: %.2f\n", num1+num2)
    case "pengurangan":
        fmt.Printf("Hasil: %.2f\n", num1-num2)
    case "perkalian":
        fmt.Printf("Hasil: %.2f\n", num1*num2)
    case "pembagian":
        if num2 != 0 {
            fmt.Printf("Hasil: %.2f\n", num1/num2)
        } else {
            fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan")
        }
    default:
        fmt.Println("Operator tidak valid")
    }
}
