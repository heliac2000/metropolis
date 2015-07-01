package util

// Create 2-dim array
//
func Create2DimArray(r, c int, cap ...int) [][]float64 {
	var arr [][]float64

	if len(cap) > 0 {
		arr = make([][]float64, r, cap[0])
	} else {
		arr = make([][]float64, r)
	}
	for i := 0; i < r; i++ {
		arr[i] = make([]float64, c)
	}

	return arr
}

// Copy 2-dim array
//
func Copy2DimArray(src [][]float64) [][]float64 {
	r, c := len(src), len(src[0])
	dst := make([][]float64, r, cap(src))
	for i := 0; i < r; i++ {
		dst[i] = make([]float64, c, cap(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}
