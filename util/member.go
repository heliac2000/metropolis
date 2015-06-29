package util

// member function
//
func Member(n int, l []int) bool {
	for _, x := range l {
		if n == x {
			return true
		}
	}
	return false
}
