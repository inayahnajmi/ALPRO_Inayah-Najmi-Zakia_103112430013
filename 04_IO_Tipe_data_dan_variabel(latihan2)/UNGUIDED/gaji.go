package main

import "fmt"

func main() {
	var jamKerjaPerMinggu, upahPerJam float64

	// Input jumlah jam kerja dalam seminggu dan upah per jam
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jamKerjaPerMinggu)
	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	// Hitung jam lembur dan gaji
	jamNormal := 40.0
	jamLembur := 0.0
	if jamKerjaPerMinggu > jamNormal {
		jamLembur = jamKerjaPerMinggu - jamNormal
		jamKerjaPerMinggu = jamNormal
	}

	// Menghitung total gaji bulanan
	gajiNormal := jamKerjaPerMinggu * upahPerJam
	gajiLembur := jamLembur * upahPerJam * 1.5
	totalGajiBulanan := (gajiNormal + gajiLembur) * 4

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: %.2f\n", totalGajiBulanan)
}
