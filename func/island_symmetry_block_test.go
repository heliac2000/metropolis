package functions_test

import (
	"testing"

	. "./"
)

type testCasesIslandSymmetry struct {
	cab      []int
	expected float64
}

func TestIslandSymmetry(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv")

	testCases := []testCasesIslandSymmetry{
		// [R]
		// IslandSymmetryBlock(list(c(1),c(1),c(1)))
		// => 2
		// IslandSymmetryBlock(list(c(313, 338, 415, 261, 337, 239, 282, 292, 253),c(1),c(1)))
		// => 1
		//
		{
			cab:      []int{1},
			expected: 2,
		},
		{
			cab:      []int{313, 338, 415, 261, 337, 239, 282, 292, 253},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		actual := IslandSymmetryBlock(tc.cab)
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
