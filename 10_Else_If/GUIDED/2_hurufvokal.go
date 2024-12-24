package main
import "fmt"

func main() {
	var x string
	var huruf, vKecil, vBesar bool
	
	fmt.Print("Masukkan satu karakter: ")
	fmt.Scanln(&x)

	if len(x) != 1 {
		fmt.Println("Input harus berupa satu karakter.")
		return
	}

	char := x[0]

	huruf = (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
	
	vKecil = char == 'a' || char == 'i' || char == 'u' || char == 'e' || char == 'o'
	vBesar = char == 'A' || char == 'I' || char == 'U' || char == 'E' || char == 'O'

	if huruf && (vKecil || vBesar) {
		fmt.Println("vokal")
	} else if huruf && !(vKecil || vBesar) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}
