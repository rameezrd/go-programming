package saying

import "testing"
import "fmt"

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome to Fun to learn golang James" {
		t.Error("Got ", s, "expected", "Welcome to Fun to learn golang James")
	}
}
func ExmpleGreet() {
	fmt.Println(Greet("James"))
	//output:
	//Welcome to Fun to learn golang
}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
