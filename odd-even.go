package go_odd_even

func OddEven(angka ...int) string {
	var result string
	for i, value := range angka {
		if i != 0 {
			result = result + ","
		}
		if value%2 == 1 {
			result = result + "Ganjil"
		} else {
			result = result + "Genap"
		}

	}

	return result

}
