package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("ctx chk:\t", ctx.Err())
	fmt.Println("Goroutine chk 1:\t", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return

			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working: ", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("ctx chk2:\t", ctx.Err())
	fmt.Println("Goroutine chk:2\t", runtime.NumGoroutine())
	fmt.Println("about to cancel")
	cancel()
	fmt.Println("Context cancelled")

	time.Sleep(time.Second * 2)
	fmt.Println("ctx chk3:\t", ctx.Err())
	fmt.Println("Goroutine chk:3\t", runtime.NumGoroutine())

}
