package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime\t:%s\narchitecture: %s", runtime.GOOS, runtime.GOARCH)
}
