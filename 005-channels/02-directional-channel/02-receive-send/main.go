package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive chan
	cs := make(chan<- int) //send chan
	go func() {
		c <- 42
		c <- 43 // unlike previous example here we allow 2 numbers to sit  make(chan int, 2) but it needs same no. of fmt

	}()
	fmt.Println("----------------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
