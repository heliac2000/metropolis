package functions_test

import (
	"testing"

	. "./"
)

type testCasesIsomorphIslands struct {
	xtest1, ctest1 []int
	otest1         []float64
	xtest2, ctest2 []int
	otest2         []float64
	expected       bool
}

func TestIsomorphIslandsBlock(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesIsomorphIslands{
		{
			// [R] Isomorph.Islands.Block(list(c(10), c(3), c(4)), list(c(20), c(3), c(4)))
			//     => TRUE
			//
			xtest1: []int{10}, ctest1: []int{2}, otest1: []float64{3},
			xtest2: []int{20}, ctest2: []int{2}, otest2: []float64{3},
			expected: true,
		},
		{
			// [R] Isomorph.Islands.Block(list(c(10,20,30), c(3,4,5), c(6,7,8)), list(c(11,21,31), c(3,4,5), c(6,7,8)))
			//     => TRUE
			//
			xtest1: []int{10, 20, 30}, ctest1: []int{3, 4, 5}, otest1: []float64{6, 7, 8},
			xtest2: []int{11, 21, 31}, ctest2: []int{3, 4, 5}, otest2: []float64{6, 7, 8},
			expected: true,
		},
		{
			// [R] Isomorph.Islands.Block(list(c(10,20,30), c(3,4,5), c(6,7,8)), list(c(11,21,31), c(3,4,5), c(6,7,9)))
			//     => FALSE
			//
			xtest1: []int{10, 20, 30}, ctest1: []int{3, 4, 5}, otest1: []float64{6, 7, 8},
			xtest2: []int{11, 21, 31}, ctest2: []int{3, 4, 5}, otest2: []float64{6, 7, 9},
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := IsomorphIslandsBlock(tc.xtest1, tc.ctest1, tc.otest1, tc.xtest2, tc.ctest2, tc.otest2)
		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
