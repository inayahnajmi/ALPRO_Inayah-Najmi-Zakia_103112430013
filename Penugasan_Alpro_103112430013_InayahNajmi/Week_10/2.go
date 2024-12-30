package main

import (
	"fmt"
)

func main() {
	var temperatures [5]float64

	fmt.Println("Masukkan 5 temperatur zat radioaktif:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Temperatur ke-%d: ", i+1)
		fmt.Scan(&temperatures[i])
	}

	result := evaluateTemperatures(temperatures)
	fmt.Println("Hasil:", result)
}

func evaluateTemperatures(temps [5]float64) string {
	stableUp := true
	stableDown := true

	for i := 0; i < 4; i++ {
		if temps[i] < temps[i+1] {
			stableDown = false
		} else if temps[i] > temps[i+1] {
			stableUp = false
		} else {
			stableUp = false
			stableDown = false
		}
	}

	if stableUp {
		return "Stabil naik/turun"
	} else if stableDown {
		return "Stabil naik/turun"
	} else {
		return "Tidak stabil"
	}
}
