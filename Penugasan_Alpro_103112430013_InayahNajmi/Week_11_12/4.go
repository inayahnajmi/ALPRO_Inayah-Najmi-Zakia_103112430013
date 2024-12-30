package main

import "fmt"

func main() {
	var n, m, x, y int

	fmt.Print("Masukkan jumlah gula (n), jumlah kopi (m), kebutuhan gula per cangkir (x), kebutuhan kopi per cangkir (y): ")
	fmt.Scan(&n, &m, &x, &y)

	var cups int

	for n >= x && m >= y {
		n -= x 
		m -= y 
		cups++ 
	}

	fmt.Printf("Dapat dibuat %d cangkir kopi.\n", cups)
}
