package main

import (
    "fmt"
)

func hitungMotor(jumlahOrang int) int {
    return (jumlahOrang + 1) / 2
}

func main() {
    var jumlahOrang int
    fmt.Print("Masukkan jumlah orang yang akan touring: ")
    fmt.Scan(&jumlahOrang)
    
    jumlahMotor := hitungMotor(jumlahOrang)
    fmt.Printf("Jumlah motor yang dibutuhkan: %d\n", jumlahMotor)
}
