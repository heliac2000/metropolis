package util

// Fill slice with a value
//
func FillSlice(slc []int, val int) {
	for i := 0; i < len(slc); i++ {
		slc[i] = val
	}
}

// func FillSlice(slc, val interface{}) {
// 	sv := reflect.ValueOf(slc)
// 	if sv.Kind() != reflect.Slice {
// 		panic("fill: slc expected slice")
// 	}
//
// 	vv := reflect.ValueOf(val)
// 	if vv.Type() != sv.Type().Elem() {
// 		panic("fill: val type != slc element type")
// 	}
//
// 	for i := 0; i < sv.Len(); i++ {
// 		sv.Index(i).Set(vv)
// 	}
// }
