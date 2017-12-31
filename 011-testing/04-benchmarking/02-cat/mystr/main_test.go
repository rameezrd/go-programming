package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shaken not stirred" {
		t.Error("Got", s, "Want", "Shaken not stirred")
	}
}
func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("Got", s, "Want", "Shaken not stirred")
	}
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	//Output:
	//Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	//OutPut:
	// Shaken not stirred
}

const s = "Whoever we have been this year, let's commit to ourselves to stand a bit taller, be a bit more noble, act a bit more wisely, be a bit more responsible and be a bit more forgiving next year.If enough of us do that, a better path will open up for all of us 2018: All hands on deck "

var xs []string

func BenchmarkCat(b *testing.B) {
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarJoint(b *testing.B) {
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
