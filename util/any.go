package util

// Any
//
func Any(arr []bool) bool {
	for _, v := range arr {
		if v {
			return true
		}
	}

	return false
}
