package main

import (
    "fmt"
)

func main() {
    var alas, tinggi int

    // Meminta input panjang alas dan tinggi segitiga
    fmt.Print("Masukkan panjang alas segitiga: ")
    fmt.Scanln(&alas)

    fmt.Print("Masukkan tinggi segitiga: ")
    fmt.Scanln(&tinggi)

    // Menghitung luas segitiga
    luas := (alas * tinggi) / 2

    // Menampilkan hasil perhitungan
    fmt.Printf("Luas segitiga dengan alas %d dan tinggi %d adalah: %d\n", alas, tinggi, luas)
}
