package saying

import (
	"fmt"
)

func Greet(s string) string {
	return fmt.Sprint("Welcome to Fun to learn golang ", s)
}
