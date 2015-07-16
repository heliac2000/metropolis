package functions_test

import (
	"testing"

	. "../util"
	. "./"
)

type testCasesBrokenIsland struct {
	Xtest    [][]int
	expected bool
}

var AdjCuml [][][]int = LoadFromCsvFileList("./data/AdjCuml.csv")

func TestSurrAdj(t *testing.T) {
	testCases := []testCasesBrokenIsland{
		{
			Xtest:    [][]int{{1, 10}, {20, 30}, {40, 50}},
			expected: true,
		},
		{
			Xtest:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: false,
		},
	}

	for _, tc := range testCases {
		broken := BrokenIslandUnitCell(tc.Xtest, AdjCuml)
		if broken != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", broken, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
