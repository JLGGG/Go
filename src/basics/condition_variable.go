package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // All cores usage.

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex) //Create a condition variable using mutex.

	c := make(chan bool, 3) //Create an asynchronous channel.

	for i:=0; i<3; i++{
		go func (n int) { //Generate three goroutines.
			mutex.Lock() //Lock mutex, starting cond.Wait() protection.
			c <- true //Send a true to channel c.
			fmt.Println("wait begin : ", n)
			cond.Wait() //Condition variable wait.
			fmt.Println("Wait end : ", n)
			mutex.Unlock() //Unlock mutex, end cond.Wait() protection.
		}(i)
	}

	for i:=0; i<3; i++ {
		<-c //Wait for when executing all goroutines.
			//If this "for instruction" don't be there, this program makes deadlock.
	}

	for i:=0; i<3; i++{
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal() //Waiting goroutines wakes up one by one.
		mutex.Unlock()
	}
	fmt.Scanln()
}
