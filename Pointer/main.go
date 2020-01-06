package main

import "fmt"

func main() {
	var numberA int = 18
	var numberB *int = &numberA // -> Variabel untuk mencari alamat memori dari variabel numberA

	fmt.Println("numberA (value):", numberA)
	fmt.Println("numberA (address):", &numberA) // -> Alamat memori dari variabel numberA

	fmt.Println("numberB (value):", *numberB) // -> Deference alamat memori dari variabel numberB
	fmt.Println("numberB (address):", numberB)
}
