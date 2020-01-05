package main

import "fmt"

func main() {
	// Operator Aritmatika
	var hasil = 1 + 3 - 1*2/2
	fmt.Println("Hasil =", hasil)

	// Operator Perbandingan
	var jumlah = 1 + 1
	var samaDengan = jumlah == 2
	fmt.Printf("Apakah %d sama dengan %d (%t) \n", jumlah, jumlah, samaDengan)

	// Operator Logika
	var trueAndTrue = true && true
	fmt.Printf("true && true = %t \n", trueAndTrue)

	var trueOrTrue = true || true
	fmt.Printf("true || true = %t \n", trueOrTrue)

	var trueAndFalse = true && false
	fmt.Printf("true && false = %t \n", trueAndFalse)

	var notTrue = !true
	fmt.Printf("!true = %t \n", notTrue)
}
