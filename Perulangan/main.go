package main

import "fmt"

func main() {
	// Perulangan angka 1 - 5
	for i := 1; i <= 5; i++ {
		fmt.Printf("Angka ke %d\n", i)
	}

	// Perulangan bersarang
	for x := 0; x < 5; x++ {
		for y := x; y < 5; y++ {
			fmt.Print(y, " ")
		}
		fmt.Println()
	}
}
