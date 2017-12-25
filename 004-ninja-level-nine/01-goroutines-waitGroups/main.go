package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPUs ", runtime.NumCPU())
	fmt.Println("Begin gs ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from thing 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from thing 2")
		wg.Done()
	}()
	fmt.Println("mid CPUs ", runtime.NumCPU())
	fmt.Println("mid gs ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("end CPUs ", runtime.NumCPU())
	fmt.Println("end gs ", runtime.NumGoroutine())

}
