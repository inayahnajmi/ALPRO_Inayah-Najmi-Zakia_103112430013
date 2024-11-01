package main

import (
    "fmt"
)

func cekFaktor(x, y int) (bool, bool) {
    var faktorX, faktorY bool
    
    if y % x == 0 {
        faktorX = true
    } else {
        faktorX = false
    }
    
    if x % y == 0 {
        faktorY = true
    } else {
        faktorY = false
    }
    
    return faktorX, faktorY
}

func main() {
    var x, y int
    fmt.Print("Masukkan bilangan x: ")
    fmt.Scan(&x)
    fmt.Print("Masukkan bilangan y: ")
    fmt.Scan(&y)
    
    faktorX, faktorY := cekFaktor(x, y)
    fmt.Printf("x adalah faktor dari y: %t\n", faktorX)
    fmt.Printf("y adalah faktor dari x: %t\n", faktorY)
}
