package main

import "fmt"

func main() {
	// Slice
	var fruits = []string{"Apel", "Melon", "Pisang", "Alpukat"}
	fmt.Println(fruits)
	fmt.Println(fruits[2])
	fmt.Println(len(fruits))
	var newFruits = append(fruits, "Jeruk")
	fmt.Println(newFruits)
}
