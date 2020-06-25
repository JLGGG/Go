package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //Using all cores of CPU.

	var data = []int{} //Create a slice of int type.

	//Create a race condition by simultaneously accessing data from a critical section through two goroutines.
	//Add 1 to data slice 1000 times in the goroutine.
	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)

			runtime.Gosched() // Yield for other goroutine to use core.
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second) // Wait for 2 second.

	fmt.Println(len(data)) // Print a length of the data slice.
}
