package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	go func() {
		c <- 42
		//c <- 43 ==> if make(chan int, 1)allows to sit only 1 number then c<-43 line will give error

	}()
	fmt.Println(<-c)
}
