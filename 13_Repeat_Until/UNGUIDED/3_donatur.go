package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	totalDonasi := 0
	jumlahDonatur := 0

	fmt.Println("Masukkan jumlah donasi tiap donatur (akhiri setelah target tercapai):")
	for totalDonasi < target {
		var donasi int
		fmt.Scan(&donasi)
		jumlahDonatur++
		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
