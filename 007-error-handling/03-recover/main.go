package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("returning normally from f.")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f ", r)
		}
	}()
	fmt.Println("calling g.")
	g(0)
	fmt.Println("returning normally fron g. ")
}
func g(i int) {
	if i > 3 {
		fmt.Println("panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("deferred in g", i)
	fmt.Println("printing in g ", i)
	g(i + 1)
}
