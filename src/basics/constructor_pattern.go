package basics

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func Constructor(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func main() {
	rect := Constructor(20, 10)

	fmt.Println(rect)
}
