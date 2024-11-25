package main

import (
	"fmt"
	"math"
)

func main() {
	var kendaraan string
	var durasi float64

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scanln(&kendaraan)
	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scanln(&durasi)

	var tarifPerJam int
	switch kendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	jamParkir := int(math.Ceil(durasi))

	totalBiaya := tarifPerJam * jamParkir

	fmt.Printf("Total biaya parkir untuk %s selama %d jam adalah Rp %d\n", kendaraan, jamParkir, totalBiaya)
}
