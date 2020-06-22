package main

import(
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)//Create synchronous channel
	count := 3

	go func() {
		for i:=0; i<count; i++{
			done <- true //Send true to goroutine and then waiting until taking out value.
			fmt.Println("Goroutine : ", i)
			time.Sleep(1*time.Second) //Wait for 1 second.
		}
	}()

	for i:=0; i<count; i++{
		<-done //Waiting until coming in value in channel, take out value.
		fmt.Println("Main function : ", i)
	}
}