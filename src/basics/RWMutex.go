package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Using all cores.

	var data = 0
	var rwMutex = new(sync.RWMutex) // Create read, write mutex.

	//Read lock don't stop each other. But, write lock is blocked because
	//the value cannot be changed during a read attempt.
	//Block both read and write locks because the previous values should not be read elsewhere
	//during a write attempt, or changed elsewhere.
	go func() {
		for i := 0; i < 3; i++ {
			//Lock a write mutex.
			rwMutex.Lock() // Starting write protection.
			data += 1
			fmt.Println("write  : ", data)
			time.Sleep(10 * time.Millisecond) //Wait for 10 millisecond.
			//Unlock the write mutex.
			rwMutex.Unlock() // End write protection.
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			//Lock a read mutex.
			rwMutex.RLock() // Starting read protection.
			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second)
			//Unlock the read mutex.
			rwMutex.RUnlock() // End read protection.
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			//Lock a read mutex.
			rwMutex.RLock()
			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second)
			//Unlock the read mutex.
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
