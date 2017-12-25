package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
			fmt.Println(incrementer)
		}()
	}
	wg.Wait()
	fmt.Println("end value increm:", incrementer)
}
