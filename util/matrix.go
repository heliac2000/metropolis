package util

import "sync"

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

// Multiply matrix A and B
//
func MatrixMultiply(A, B [][]int) [][]int {
	r_A, c_A, c_B := len(A), len(A[0]), len(B[0])

	C := Create2DimArray(int(0), r_A, c_B).([][]int)

	var wg sync.WaitGroup
	wg.Add(r_A)
	for i := 0; i < r_A; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < c_A; j++ {
				for k := 0; k < c_A; k++ {
					C[i][j] += A[i][k] * B[k][j]
				}
			}
		}(i)
	}
	wg.Wait()

	return C
}

// Add B to A
//
func MatrixAdd(A, B [][]int) [][]int {
	r, c := len(A), len(A[0])

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			A[i][j] += B[i][j]
		}
	}

	return A
}
