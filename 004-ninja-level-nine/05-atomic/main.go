package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var incrementer int64
	var wg sync.WaitGroup
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println("Increm:\t", atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value increm:", incrementer)
}
