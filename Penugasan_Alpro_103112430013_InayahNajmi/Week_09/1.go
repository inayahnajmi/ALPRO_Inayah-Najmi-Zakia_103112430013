package main

import (
	"fmt"
)

func main() {
	var n int
	const kapasitasMobil = 7

	fmt.Print("Masukkan jumlah mahasiswa (N): ")
	fmt.Scan(&n)

	jumlahMobil := n / kapasitasMobil
	if n%kapasitasMobil != 0 {
		jumlahMobil++
	}

	fmt.Println("Jumlah mobil yang diperlukan:", jumlahMobil)
}