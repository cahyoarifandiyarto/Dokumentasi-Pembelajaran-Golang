package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("Cahyo Arif Andiyarto")
	go sayHelloTo("Ilham")
	go sayHelloTo("Akbar")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	// Proses nya
	// Variabel data mengirim ke channel (Variabel messages)
	// Lalu variabel message1, message2, message3 sebeagai penerima dari channel variabel message

	// Channel sebagai tipe data parameter
	for _, each := range []string{"Cahyo", "Arif", "Andiyarto"} {
		go func(who string) {
			var data = fmt.Sprintf("Hai %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

// Channel sebagai tipe data parameter
func printMessage(what chan string) {
	fmt.Println(<-what)
}
