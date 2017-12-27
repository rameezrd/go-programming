package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	cr := make(chan<- int)
	go func() {
		cs <- 42
		cr <- 43
	}()
	fmt.Println(<-cs)
	fmt.Println(cr)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
