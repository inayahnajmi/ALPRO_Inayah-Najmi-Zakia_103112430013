package main

import (
    "fmt"
)

func main() {
    var sisi int

    // Meminta input panjang sisi kubus
    fmt.Print("Masukkan panjang sisi kubus: ")
    fmt.Scanln(&sisi)

    // Menghitung volume kubus
    volume := sisi * sisi * sisi

    // Menampilkan hasil perhitungan
    fmt.Printf("Volume kubus dengan panjang sisi %d adalah: %d\n", sisi, volume)
}
