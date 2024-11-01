package main

import (
    "fmt"
)

func cekGenapNegatif(bilangan int) string {
    if bilangan < 0 && bilangan%2 == 0 {
        return "genap negatif"
    }
    return "bukan"
}

func main() {
    var bilangan int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)
    
    hasil := cekGenapNegatif(bilangan)
    fmt.Printf(hasil)
}
