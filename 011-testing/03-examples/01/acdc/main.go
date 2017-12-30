//Package acdc asks if you are ready to rock
package acdc

//Sum adds unlimited number of vlue of the type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
