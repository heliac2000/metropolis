package util

import (
	"reflect"
	"sync"
)

// Transpose
//
func Transpose(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	r, c := v.Len(), v.Index(0).Len()
	if r == 0 || c == 0 {
		return m
	}

	ts := reflect.SliceOf(v.Index(0).Index(0).Type())
	ts2 := reflect.SliceOf(ts)

	t := reflect.MakeSlice(ts2, c, c)
	for i := 0; i < c; i++ {
		t.Index(i).Set(reflect.MakeSlice(ts, r, r))
		for j := 0; j < r; j++ {
			t.Index(i).Index(j).Set(v.Index(j).Index(i))
		}
	}

	return t.Interface()
}

// Multiply matrix A and B
//
func MatrixMultiply(A, B [][]int) [][]int {
	var C [][]int
	r_A, c_A, c_B := len(A), len(A[0]), len(B[0])

	Create2DimArray(&C, r_A, c_B)

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
