package main

import (
	"fmt"

	"github.com/FunToLearn/go-programming/010-ninja-level-12/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
