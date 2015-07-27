package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testReductionBlock struct {
	xtest, ctest, otest []int
	expected            [][][]int
}

func TestReductionBlock(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	testCases := []testReductionBlock{
		{
			xtest:    []int{1},
			ctest:    []int{2},
			otest:    []int{3},
			expected: [][][]int{{{0, 0, 0}}},
		},
		{
			// [R] writeListData(Reduction.Block(list(1:10,11:20,21:30)), "ReductionBlock_01.csv")
			//
			xtest:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			ctest:    []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			otest:    []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			expected: LoadFromCsvFileList("./data/ReductionBlock_01.csv"),
		},
	}

	for _, tc := range testCases {
		actual := ReductionBlock(tc.xtest, tc.ctest, tc.otest)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
