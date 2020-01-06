package main

import "fmt"

func main() {
	// Array
	var names [3]string
	names[0] = "Cahyo"
	names[1] = "Arif"
	names[2] = "Andiyarto"

	fmt.Println(names[0], names[1], names[2])

	// Inisialisasi nilai awal pada array
	var fruits = [4]string{"melon", "apel", "alpukat", "jeruk"}

	fmt.Println("Panjang array variabel fruits", len(fruits))
	fmt.Println("Nilai array pada variabel fruits", fruits)

	// Inisialisasi nilai array tanpa jumlah elemen
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Panjang array variabel numbers", len(numbers))
	fmt.Println("Nilai array pada variabel numbers", numbers)

	// Array multidimensi
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{[3]int{1, 4, 2}, [3]int{5, 2, 1}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)

	// Perulangan array dengan for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%d %s\n", i, fruits[i])
	}

	// Perulangan array dengan for - range
	for j, fruit := range fruits {
		fmt.Printf("%d %s\n", j, fruit)
	}

	// Penggunaan variabel underscore dalam for - range
	for _, fruitt := range fruits {
		fmt.Printf("%s\n", fruitt)
	}
}
