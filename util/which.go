package util

// Count a number of items in slice
//
func CountItems(item int, slice []int) int {
	n := 0
	for _, v := range slice {
		if v == item {
			n++
		}
	}

	return n
}

// R's which function
//
func Which(slice []int, a int) []int {
	c := make([]int, 0, len(slice))
	for i, v := range slice {
		if v == a {
			c = append(c, i)
		}
	}

	return c
}

// R's `which' functions
//
// func WhichOverZero(hop int, ahop [][]int) []int {
// 	l := len(ahop[hop])
// 	hop_d := make([]int, 0, l)
// 	for i := 0; i < l; i++ {
// 		if ahop[hop][i] > 0 {
// 			hop_d = append(hop_d, i)
// 		}
// 	}
//
// 	return hop_d
// }

func WhichIn(a, b []int) bool {

	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				return true
			}
		}
	}

	return false
}

func WhichNotIn(a, b []int) []int {
	c := make([]int, 0, len(a))

	for _, v1 := range a {
		if !Member(v1, b) {
			c = append(c, v1)
		}
	}

	return c
}
