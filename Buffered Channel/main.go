package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Buffered channel
	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("Menerima data:", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Mengirim data", i)
		messages <- i
	}
}
