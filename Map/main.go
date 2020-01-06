package main

import "fmt"

func main() {
	// Map
	var biodata map[string]string
	biodata = map[string]string{}

	biodata["Nama"] = "Cahyo Arif Andiyarto"
	biodata["Umur"] = "18"
	biodata["Domisili"] = "Sleman"

	fmt.Println("Nama", biodata["Nama"])
	fmt.Println("Umur", biodata["Umur"])
	fmt.Println("Domisili", biodata["Domisili"])

	// Inisialisasi nilai map
	var data map[string]int
	data = map[string]int{}
	data["satu"] = 1

	fmt.Println("Data ke ->", data["satu"])

	// Looping map menggunakan for - range
	var months = map[string]int{
		"januari":  01,
		"februari": 02,
		"maret":    03,
		"april":    04,
	}

	for key, val := range months {
		fmt.Println("Bulan", key, "Adalah bulan ke ->", val)
	}

	fmt.Println("Jumlah key pada variabel months ->", len(months))
	fmt.Println("Data pada variabel months ->", months)
}
