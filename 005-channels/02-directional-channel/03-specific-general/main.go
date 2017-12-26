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
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//c = cr (if we use this way which is generl to specific it will give error)
	//c = cs (if we use this way which is generl to specific it will give error)

	cr = c //specific to general
	cs = c //specific to general

}
