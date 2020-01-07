package main

import "fmt"

func main() {
	var secret interface{}

	secret = "Cahyo Arif Andiyarto"
	fmt.Println(secret)

	secret = []string{"Apel", "Jeruk", "Melon"}
	fmt.Println(secret)

	secret = 18.3
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"Nama":                   "Cahyo Arif Andiyarto",
		"Umur":                   18,
		"Bahasa program disukai": []string{"Golang", "Java", "NodeJS"},
	}
	fmt.Println(data)
}
