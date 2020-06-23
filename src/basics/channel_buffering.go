package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // Create an asynchronous channel with two buffers
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true // Send true to channel, wait when buffer is full.
			fmt.Println("Goroutine : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done // Wait if buffer has no value, eject value.
		fmt.Println("Main func : ", i)
	}
}
