package functions_test

import (
	"testing"

	. "./"
)

type testCasesBrokenIsland struct {
	xtest    []int
	expected bool
}

func TestBrokenIslandUnitCell(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	testCases := []testCasesBrokenIsland{
		{
			xtest:    []int{1, 10, 20, 30, 40, 50},
			expected: true,
		},
		{
			xtest:    []int{1, 2, 3, 4, 5, 6},
			expected: false,
		},
	}

	for _, tc := range testCases {
		broken := BrokenIslandUnitCell(tc.xtest)
		if broken != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", broken, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
