package util

import (
	"math"
	"sort"
	"sync"

	"github.com/skelterjohn/go.matrix"
)

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

	C := Create2DimArrayInt(r_A, c_B)

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

func MatrixMultiplyFloat(A, B [][]float64) [][]float64 {
	r_A, c_A, c_B := len(A), len(A[0]), len(B[0])

	C := Create2DimArrayFloat(r_A, c_B)

	var wg sync.WaitGroup
	wg.Add(r_A)
	for i := 0; i < r_A; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < c_B; j++ {
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

// [R] dist function
//
func Dist(m [][]float64, diag float64) [][]float64 {
	r, c := len(m), len(m[0])
	dist := Create2DimArrayFloat(r, r)

	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			if i == j {
				dist[i][j] = diag
			} else {
				sum := 0.0
				for k := 0; k < c; k++ {
					sum += (m[i][k] - m[j][k]) * (m[i][k] - m[j][k])
				}
				dist[i][j] = math.Sqrt(sum)
			}
		}
	}

	return dist
}

// [R] eigen function(return eigen$values only)
//
func EigenValues(mat [][]float64) []float64 {
	l := len(mat) // Square matrix
	m := make([]float64, l*l)
	for i, k := 0, 0; i < l; i++ {
		for j := 0; j < l; j++ {
			m[k], k = mat[i][j], k+1
		}
	}

	dm := matrix.MakeDenseMatrix(m, l, l)
	_, v, err := dm.Eigen()
	if err != nil {
		return m
	}

	ev := make([]float64, l)
	for i := 0; i < l; i++ {
		ev[i] = v.Get(i, i)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(ev)))

	return ev
}

// [R] Krls gausskernel function
//
func GaussKernel(x [][]float64, sigma float64) [][]float64 {
	xd := Dist(x, 0.0)
	for i := 0; i < len(xd); i++ {
		for j := 0; j < len(xd[0]); j++ {
			xd[i][j] = math.Exp(-1 * (xd[i][j] * xd[i][j]) / sigma)
		}
	}

	return xd
}

// Transpose
//
// func Transpose(m interface{}) interface{} {
// 	v := reflect.ValueOf(m)
// 	r, c := v.Len(), v.Index(0).Len()
// 	if r == 0 || c == 0 {
// 		return m
// 	}
//
// 	ts := reflect.SliceOf(v.Index(0).Index(0).Type())
// 	ts2 := reflect.SliceOf(ts)
//
// 	t := reflect.MakeSlice(ts2, c, c)
// 	for i := 0; i < c; i++ {
// 		t.Index(i).Set(reflect.MakeSlice(ts, r, r))
// 		for j := 0; j < r; j++ {
// 			t.Index(i).Index(j).Set(v.Index(j).Index(i))
// 		}
// 	}
//
// 	return t.Interface()
// }
