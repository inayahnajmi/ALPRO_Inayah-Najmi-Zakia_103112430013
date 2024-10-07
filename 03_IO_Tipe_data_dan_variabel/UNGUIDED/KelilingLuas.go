package main

import (
    "fmt"
)

func main() {
    // Panjang sisi alun-alun dalam meter
    var sisi float64 = 27

    // Menghitung keliling dan luas
    keliling := 4 * sisi
    luas := sisi * sisi

    // Menampilkan hasil
    fmt.Println("Keliling Alun-alun Purwokerto adalah:", keliling, "meter")
    fmt.Println("Luas Alun-alun Purwokerto adalah:", luas, "meter persegi")
}
