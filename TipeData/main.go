package main

import "fmt"

func main() {
	// Tipe Data Numerik Non Desimal
	var positiveNumber uint8 = 200
	negativeNumber := -1234121412

	fmt.Printf("Bilangan positif: %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif: %d\n", negativeNumber)

	// Tipe Data Numerik Desimal
	var decimalNumber = 14.500
	var price = 20.500
	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f\n", decimalNumber)
	fmt.Printf("Harga: %f\n", price)

	// Tipe Data Boolean
	var isMale = true
	fmt.Printf("Apakah dia laki-laki? %t \n", isMale)

	// Tipe Data String
	var str = "Ini adalah tipe data string"
	fmt.Println(str)
}
