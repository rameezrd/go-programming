package main

import (
	"fmt"

	"github.com/rameezrd/go-programming/009-documentation/01/mymath"
)

func main() {
	fmt.Println("2 + 3 =\t", mymath.Sum{2, 3})
	fmt.Println("4 + 7 =\t", mymath.Sum{4, 7})
	fmt.Println("5 + 9 =\t", mymath.Sum{5, 9})
}
