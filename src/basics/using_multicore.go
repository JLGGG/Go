package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Set the maximum number of cores to use after obtaining the number of cores.

	fmt.Println(runtime.GOMAXPROCS(0)) // Print setting value.

	s := "Hello, world"

	for i := 0; i < 100; i++ {
		go func(n int) { // Run anonymous function as goroutine.
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
