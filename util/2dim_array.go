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
func Copy2DimArray(src interface{}) interface{} {
	v := reflect.ValueOf(src)
	r, c := v.Len(), v.Index(0).Len()
	ts := reflect.SliceOf(v.Index(0).Index(0).Type())
	ts2 := reflect.SliceOf(ts)

	dst := reflect.MakeSlice(ts2, r, r)
	for i := 0; i < r; i++ {
		dst.Index(i).Set(reflect.MakeSlice(ts, c, c))
		reflect.Copy(dst.Index(i), v.Index(i))
	}

	return dst.Interface()
}

// Copy vector
//
func CopyVector(src interface{}) interface{} {
	v := reflect.ValueOf(src)
	dst := reflect.MakeSlice(v.Type(), v.Len(), v.Len())

	reflect.Copy(dst, v)

	return dst.Interface()
}
