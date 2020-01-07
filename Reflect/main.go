package main

import (
	"fmt"
	"reflect"
)

func main() {
	var age = 18
	var reflectValue = reflect.ValueOf(age)

	fmt.Println("Tipe variabel age adalah:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel age adalah:", reflectValue.Int())
	}
}
