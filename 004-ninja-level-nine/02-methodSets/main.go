package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
	}
	saySomething(&p1)
	//or
	p1.speak()

	//fmt.Println()
}
