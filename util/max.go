package util

// Return the biggest value in a slice of ints.
func Max(slice ...int) int {
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}

	return max
}

// Return the smallest value in a slice of ints.
func Min(slice ...int) int {
	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}

	return min
}
