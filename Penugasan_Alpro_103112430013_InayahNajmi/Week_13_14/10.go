package main

import "fmt"

func main() {
    var x_0013, baris, kolom int
    fmt.Scan(&x_0013)

    for baris = 1; baris <= x_0013; baris++ {
        for kolom = 1; kolom <= x_0013; kolom++ {
            if (baris == kolom) || (kolom+baris-1 == x_0013) {
                fmt.Print(baris, " ")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }      
} 