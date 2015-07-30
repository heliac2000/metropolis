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
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")
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
			xout:  nthColumn(0, out),
			cout:  nthColumn(1, out),
			oout:  nthColumnFloat(2, out),
		},
	}

	for _, tc := range testCases {
		xout, cout, oout := ReductionBlock(tc.xtest, tc.ctest, tc.otest)
		if !reflect.DeepEqual(xout, tc.xout) ||
			!reflect.DeepEqual(cout, tc.cout) ||
			!reflect.DeepEqual(oout, tc.oout) {
			t.Errorf("\ngot  %v\nwant %v", xout, tc.xout)
			return
		}
	}
}

func nthColumn(n int, data [][][]int) [][]int {
	r, c := len(data), len(data[0])
	ret := Create2DimArray(int(0), r, c).([][]int)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ret[i][j] = data[i][j][n]
		}
	}

	return ret
}

func nthColumnFloat(n int, data [][][]int) [][]float64 {
	r, c := len(data), len(data[0])
	ret := Create2DimArray(float64(0), r, c).([][]float64)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ret[i][j] = float64(data[i][j][n])
		}
	}

	return ret
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
