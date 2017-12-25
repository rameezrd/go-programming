package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	go func() {
		c <- 42
		c <- 43 // unlike previous example here we allow 2 numbers to sit  make(chan int, 2) but it needs same no. of fmt

	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}
