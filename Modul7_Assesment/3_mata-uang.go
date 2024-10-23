package main

import (
	"fmt"
)

func main() {
	var qirat int

	fmt.Print("Masukkan jumlah uang dalam satuan qirat: ")
	fmt.Scanln(&qirat)

	const (
		qiratPerFals  = 6
		falsPerDirham = 10
		dirhamPerDinar = 10
	)

	totalFals := qirat / qiratPerFals
	totalDirham := totalFals / falsPerDirham
	totalDinar := totalDirham / dirhamPerDinar

	sisaFals := totalFals % falsPerDirham
	sisaDirham := totalDirham % dirhamPerDinar

	fmt.Printf("Jumlah uang dalam satuan:\n")
	fmt.Printf("Dinar: %d\n", totalDinar)
	fmt.Printf("Dirham: %d\n", sisaDirham)
	fmt.Printf("Fals: %d\n", sisaFals)
	fmt.Printf("Qirat: %d\n", qirat%qiratPerFals)
}