//Package mymath providesAcme Inc math solution
package mymath

//Sum adds an unlimited number of value of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
