package basics

import "fmt"

func calc() func(x int) int {
	//The local variable is extinguished when the function is finished,
	a, b := 3, 5
	//However, local variables a,b can be used even after the function ends by using a closure.
	return func(x int) int {
		return a*x + b
	}
}

func main() {
	//Call function to store returned closure in f.
	f := calc()

	fmt.Println(f(1))
	fmt.Println(f(2))
}
