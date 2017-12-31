package main

import (
	"fmt"

	"github.com/FunToLearn/go-programming/012-ninja-level-13/02/quote"
	"github.com/FunToLearn/go-programming/012-ninja-level-13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range UseCount(quote.SunAlso) {
		fmt.Println(k, v)
	}
}
