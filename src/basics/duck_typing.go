package main

import "fmt"

type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("Quack")
}

func (d Duck) feathers() {
	fmt.Println("Duck has white and gray feathers.")
}

type Person struct {
}

func (p Person) quack() {
	fmt.Println("A man mimics a duck.")
}
func (p Person) feathers() {
	fmt.Println("People show the feathers of birds.")
}

//Define a Quacker interface that has quack, feathers method.
type Quacker interface {
	quack()
	feathers()
}

func caller(q Quacker) {
	q.quack()
	q.feathers()
}
func main() {
	var duck Duck
	var james Person

	caller(duck)  // Call quack, feather method of duck through interface.
	caller(james) // Call quack, feather method of person through interface.
}
