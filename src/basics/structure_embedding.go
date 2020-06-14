package basics

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) greeting() {
	fmt.Println("Hello")
}

type Student struct {
	Person //is-a relationship
	school string
	grade int
}
//Overriding
func (p *Student) greeting() {
	fmt.Println("Hello student")
}
func main() {
	var s Student

	s.Person.greeting()
	s.greeting()
}