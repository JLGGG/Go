package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Using all cores of CPU.

	var data = []int{}
	var mutex = new(sync.Mutex) // Create a mutex.

	//Create a race condition by simultaneously accessing data from a critical section through two goroutines.
	//Add 1 to data slice 1000 times in the goroutine.
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()           //Starting data slice protection.
			data = append(data, 1) //Add 1 to data slice.
			mutex.Unlock()         //End data slice protection.

			runtime.Gosched() //Yield for other goroutine to use core.
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()           //Starting data slice protection.
			data = append(data, 1) //Add 1 to data slice.
			mutex.Unlock()         //End data slice protection.

			runtime.Gosched() //Yield for other goroutine to use core.
		}
	}()

	time.Sleep(2 * time.Second) //Wait for 2 second.

	fmt.Println(len(data))
}
