package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testReductionBlock struct {
	xtest, ctest []int
	otest        []float64
	xout, cout   [][]int
	oout         [][]float64
}

func TestReductionBlock(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv")
	out := LoadFromCsvFileList("./data/ReductionBlock_01.csv")

	testCases := []testReductionBlock{
		{
			xtest: []int{1},
			ctest: []int{2},
			otest: []float64{3},
			xout:  [][]int{{0}},
			cout:  [][]int{{0}},
			oout:  [][]float64{{0}},
		},
		{
			// [R] writeListData(Reduction.Block(list(1:10,11:20,21:30)), "ReductionBlock_01.csv")
			//
			xtest: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			ctest: []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			otest: []float64{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			xout:  nthColumn(0, out, int(0)).([][]int),
			cout:  nthColumn(1, out, int(0)).([][]int),
			oout:  nthColumn(2, out, float64(0)).([][]float64),
		},
	}

	for _, tc := range testCases {
		xout, cout, oout := ReductionBlock(tc.xtest, tc.ctest, tc.otest)
		if !reflect.DeepEqual(xout, tc.xout) {
			t.Errorf("xtest:\ngot  %v\nwant %v", xout, tc.xout)
		} else if !reflect.DeepEqual(cout, tc.cout) {
			t.Errorf("ctest:\ngot  %v\nwant %v", cout, tc.cout)
		} else if !reflect.DeepEqual(oout, tc.oout) {
			t.Errorf("otest:\ngot  %v\nwant %v", oout, tc.oout)
		}
	}
}

func nthColumn(n int, data [][][]int, x interface{}) interface{} {
	r, c := len(data), len(data[0])

	ts := reflect.SliceOf(reflect.ValueOf(x).Type())
	ts2 := reflect.SliceOf(ts)
	arr := reflect.MakeSlice(ts2, r, r)
	for i := 0; i < r; i++ {
		arr.Index(i).Set(reflect.MakeSlice(ts, c, c))
	}

	switch x.(type) {
	case float32, float64:
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				arr.Index(i).Index(j).Set(reflect.ValueOf(float64(data[i][j][n])))
			}
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				arr.Index(i).Index(j).Set(reflect.ValueOf(data[i][j][n]))
			}
		}
	}

	return arr.Interface()
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
