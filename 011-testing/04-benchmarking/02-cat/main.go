package main

import (
	"fmt"
	"strings"

	"github.com/FunToLearn/go-programming/011-testing/04-benchmarking/02-cat/mystr"
)

const s = "Whoever we have been this year, let's commit to ourselves to stand a bit taller, be a bit more noble, act a bit more wisely, be a bit more responsible and be a bit more forgiving next year.If enough of us do that, a better path will open up for all of us 2018: All hands on deck"

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
