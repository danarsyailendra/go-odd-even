package go_odd_even

import "fmt"

func OddEven(angka int) {
	if angka%2 == 1 {
		fmt.Println("Ganjil")
	} else {
		fmt.Println("Genap")
	}

}
