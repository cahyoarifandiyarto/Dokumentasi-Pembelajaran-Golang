package main

import "fmt"

func main() {
	// Manifest Typing
	var firstName string = "Cahyo"
	var lastName string = "Arif Andiyarto"
	var myAge int = 18

	fmt.Println("Nama saya adalah", firstName, lastName)
	fmt.Println("Umur saya sekarang:", myAge, " Tahun")

	// Type inference
	city := "Sidoarjo"
	province := "Jawa Timur"
	postalCode := 61254
	fmt.Println("Kota:", city)
	fmt.Println("Provinsi:", province)
	fmt.Println("Kode pos:", postalCode)

	// Multi variable
	first, second, third := "Satu", "Dua", "Tiga"
	fmt.Println(first, second, third)
}
