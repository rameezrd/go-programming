// Package dog allows us to more fully provides functionality for understanding dogs
package dog

//Years convert human years to dog years.
func Years(n int) int {
	return n * 10
}

//Years convert human years to dog years.
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
