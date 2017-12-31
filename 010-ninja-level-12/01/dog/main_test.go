package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(7)
	if n != 70 {
		t.Error("Got", n, "Want", 70)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	//Output:
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("Got", n, "Want", 70)
	}

}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
