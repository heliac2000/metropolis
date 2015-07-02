package util

import "reflect"

// Create 2-dim array
//
func Create2DimArray(t interface{}, r, c int, cap ...int) interface{} {
	var arr reflect.Value

	ts := reflect.SliceOf(reflect.ValueOf(t).Type())
	ts2 := reflect.SliceOf(ts)
	if len(cap) > 0 {
		arr = reflect.MakeSlice(ts2, r, cap[0])
	} else {
		arr = reflect.MakeSlice(ts2, r, r)
	}

	for i := 0; i < r; i++ {
		arr.Index(i).Set(reflect.MakeSlice(ts, c, c))
	}

	return arr.Interface()
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
