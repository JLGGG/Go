package basics

import "fmt"

type Rectangle struct {
	width  int
	height int
}
//Connect a Rectangle structure using a receiver pointer variable.
func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

func main() {
	rect := Rectangle{10, 20}

	fmt.Println(rect.area())
}
