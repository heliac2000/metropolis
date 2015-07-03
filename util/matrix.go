package util

// Transpose
//
func Transpose(m [][]float64) [][]float64 {
	r, c := len(m), len(m[0])
	if r == 0 || c == 0 {
		return m
	}

	t := make([][]float64, c)
	for i := 0; i < c; i++ {
		t[i] = make([]float64, r)
		for j := 0; j < r; j++ {
			t[i][j] = m[j][i]
		}
	}

	return t
}
