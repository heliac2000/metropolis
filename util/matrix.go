package util

import "sync"

// Create 2-dim array
//
func Create2DimArray(r, c int) [][]int {
	arr := make([][]int, r)
	for i := 0; i < r; i++ {
		arr[i] = make([]int, c)
	}
	return arr
}

func Create2DimArrayFloat(r, c int, cap ...int) [][]float64 {
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
func Copy2DimArray(src [][]int) [][]int {
	r, c := len(src), len(src[0])
	dst := make([][]int, r, cap(src))
	for i := 0; i < r; i++ {
		dst[i] = make([]int, c, cap(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func Copy2DimArrayFloat(src [][]float64) [][]float64 {
	r, c := len(src), len(src[0])
	dst := make([][]float64, r, cap(src))
	for i := 0; i < r; i++ {
		dst[i] = make([]float64, c, cap(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// Copy list(3-dim array)
//
func CopyList(src [][][]int) [][][]int {
	l := len(src)
	dst := make([][][]int, l, cap(src))
	for i := 0; i < l; i++ {
		dst[i] = Copy2DimArray(src[i])
	}

	return dst
}

// Multiply matrix A and B
//
func MatrixMultiply(A, B [][]int) [][]int {
	r_A, c_A, c_B := len(A), len(A[0]), len(B[0])

	C := Create2DimArray(r_A, c_B)

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

// Multiply matrix A and B with one thread
//
func MatrixMultiply1(A, B [][]int) [][]int {
	//r_A, c_A, r_B, c_B := len(A), len(A[0]), len(B), len(B[0])
	r_A, c_A, c_B := len(A), len(A[0]), len(B[0])

	C := Create2DimArray(r_A, c_B)

	for i := 0; i < r_A; i++ {
		for j := 0; j < c_A; j++ {
			for k := 0; k < c_A; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

// Intersection
//
func Intersection(u, v []int) []int {
	is := make([]int, 0, len(u))
	for _, x := range u {
		if Member(x, v) && !Member(x, is) {
			is = append(is, x)
		}
	}

	return is
}

// Transpose
//
func Transpose(m [][]int) [][]int {
	r, c := len(m), len(m[0])
	if r > 0 && c > 0 && r == c {
		for i := 0; i < r; i++ {
			for j := i + 1; j < c; j++ {
				if i != j {
					m[i][j], m[j][i] = m[j][i], m[i][j]
				}
			}
		}
	}

	return m
}

// start := time.Now()
// MatrixMultiply1(Ahop, Ahop)
// end := time.Now()
// fmt.Printf("%f ms\n", (float64)(end.Sub(start).Nanoseconds())/1000000.0)
//
// start = time.Now()
// MatrixMultiply(Ahop, Ahop)
// end = time.Now()
// fmt.Printf("%f ms\n", (float64)(end.Sub(start).Nanoseconds())/1000000.0)
