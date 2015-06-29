package util

// Append matrix Z2 to matrix Z1 (Z1 and Z2 have same number of columns)
//
func AddMatrix(Z1, Z2 [][]int) [][]int {
	lz1, lz2 := len(Z1), len(Z2)
	Z3 := Create2DimArray(lz1+lz2, len(Z1[0]))

	for i := 0; i < lz1; i++ {
		copy(Z3[i], Z1[i])
	}
	for i := 0; i < lz2; i++ {
		copy(Z3[i+lz1], Z2[i])
	}

	return Z3
}
