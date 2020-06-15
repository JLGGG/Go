package basics

import "fmt"

type INT int //Define int type as INT.

//Connect the Print method to INT.
func (i INT) Print() {
	fmt.Println(i)
}

type Rectangle struct {
	width, height int
}
//Connect the Print method to Rectangle.
func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}
//Define an interface with Print().
type Printer interface {
	Print()
}

func main() {
	var i INT = 5
	r := Rectangle{10, 20}

	pArr := []Printer{i, r} // Initialize interface in slice form.
	//Call Print method while touring slices.
	for index := range pArr {
		pArr[index].Print()
	}

	for _, value := range pArr {
		value.Print()
	}
}
