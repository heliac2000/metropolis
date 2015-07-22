package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCasesIsomorphIslands struct {
	xtest1, ctest1, otest1 []int
	xtest2, ctest2, otest2 []int
	expected               bool
}

func TestIsomorphIslandsBlock(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesIsomorphIslands{
		{
			// [R] Isomorph.Islands.Block(list(c(2), c(3), c(4)), list(c(2), c(3), c(4)))
			//     => TRUE
			//
			xtest1: []int{1}, ctest1: []int{2}, otest1: []int{3},
			xtest2: []int{1}, ctest2: []int{2}, otest2: []int{3},
			expected: true,
		},
	}

	for _, tc := range testCases {
		actual := IsomorphIslandsBlock(tc.xtest1, tc.ctest1, tc.otest1, tc.xtest2, tc.ctest2, tc.otest2)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
