package main

import (
    "fmt"
)

func main() {
    var idr float64
    const kurs float64 = 15000.0

    // Meminta input jumlah uang dalam IDR
    fmt.Print("Masukkan jumlah uang dalam IDR: ")
    fmt.Scanln(&idr)

    // Menghitung konversi IDR ke USD
    usd := idr / kurs

    // Menampilkan hasil perhitungan
    fmt.Printf("Jumlah %.2f IDR setara dengan %.2f USD\n", idr, usd)
}
