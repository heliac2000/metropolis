package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCasesRotationBlock struct {
	islandP  []int
	expected []int
}

func TestRotationBlock(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesRotationBlock{
		{
			// [R] Rotation.Block(list(c(1), c(1), c(1)))
			//     => [[1]][1] 1
			//
			islandP:  []int{0},
			expected: []int{0},
		},
		{
			// [R] Rotation.Block(list(c(1, 10, 100), c(1), c(1)))
			//     => [[1]][1] 19 10  1
			//
			islandP:  []int{0, 9, 99},
			expected: []int{18, 9, 0},
		},
		{
			// [R] Rotation.Block(list(c(5, 40, 200), c(1), c(1)))
			//     => [[1]][1] 75 40 7
			//
			islandP:  []int{4, 39, 199},
			expected: []int{74, 39, 6},
		},
	}

	for _, tc := range testCases {
		actual := RotationBlock(tc.islandP)
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
