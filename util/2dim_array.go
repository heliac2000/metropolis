package util

import "reflect"

// Create 2-dim array
//
func Create2DimArrayInt(r, c int, cap ...int) [][]int {
	var dim [][]int
	if len(cap) > 0 {
		dim = make([][]int, r, cap[0])
	} else {
		dim = make([][]int, r)
	}

	for i := 0; i < r; i++ {
		dim[i] = make([]int, c)
	}

	return dim
}

func Create2DimArrayFloat(r, c int, cap ...int) [][]float64 {
	var dim [][]float64
	if len(cap) > 0 {
		dim = make([][]float64, r, cap[0])
	} else {
		dim = make([][]float64, r)
	}

	for i := 0; i < r; i++ {
		dim[i] = make([]float64, c)
	}

	return dim
}

// Copy 2-dim array
//
func Copy2DimArray(dst, src interface{}) bool {
	sv, dv := reflect.ValueOf(src), reflect.ValueOf(dst).Elem()
	typ, r := sv.Type(), sv.Len()

	dv.Set(reflect.MakeSlice(typ, r, r))
	for i, t := 0, typ.Elem(); i < r; i++ {
		c := sv.Index(i).Len()
		dv.Index(i).Set(reflect.MakeSlice(t, c, c))
		reflect.Copy(dv.Index(i), sv.Index(i))
	}

	return true
}

// Copy vector
//
func CopyVectorInt(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)

	return dst
}

func CopyVectorFloat(src []float64) []float64 {
	dst := make([]float64, len(src))
	copy(dst, src)

	return dst
}

// Create 2-dim array
//
// func Create2DimArray(dim interface{}, r, c int, cap ...int) bool {
// 	v := reflect.ValueOf(dim).Elem()
// 	typ := v.Type()
// 	if len(cap) > 0 {
// 		v.Set(reflect.MakeSlice(typ, r, cap[0]))
// 	} else {
// 		v.Set(reflect.MakeSlice(typ, r, r))
// 	}
//
// 	for i, t := 0, typ.Elem(); i < r; i++ {
// 		v.Index(i).Set(reflect.MakeSlice(t, c, c))
// 	}
//
// 	return true
// }

// Copy vector
//
// func CopyVector(src interface{}) interface{} {
// 	v := reflect.ValueOf(src)
// 	dst := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
//
// 	reflect.Copy(dst, v)
//
// 	return dst.Interface()
// }
