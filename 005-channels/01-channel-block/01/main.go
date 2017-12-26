package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42 //send
	}()
	fmt.Println(<-c) //receive
}
