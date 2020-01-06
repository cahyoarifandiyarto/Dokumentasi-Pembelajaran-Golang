package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Memanggil fungsi
	printName("Cahyo Arif Andiyarto")

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number:", randomValue)
}

// Membuat fungsi untuk menampilkan nama kita
func printName(name string) {
	fmt.Println("Halo", name)
}

// Membuat fungsi dengan return value / nilai balik
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
