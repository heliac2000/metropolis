package functions_test

import (
	"testing"

	. "./"
)

type testCasesReactionClassId struct {
	diff     []int
	expected int
}

func TestReactionClassID(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	//
	testCases := []testCasesReactionClassId{
		{
			diff:     []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: 8, // REACT_CLASS9
		},
		{
			diff:     []int{0, -1, 0, 0, 2, 0, 0, 0, 0},
			expected: 6, // REACT_CLASS7
		},
		{
			diff:     []int{0, 3, 0, 0, 0, 0, 0, 0, 0},
			expected: -1, // REACT_UNKNOWN
		},
	}

	for _, tc := range testCases {
		actual := ReactionClassID(tc.diff)
		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
