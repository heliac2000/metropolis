package functions_test

import (
	"math"
	"testing"

	. "../util"
	. "./"
)

type testCasesShiftMCpos struct {
	X        []float64
	expected [][]float64
}

func TestShiftMCpos(t *testing.T) {
	testCases := []testCasesShiftMCpos{
		{
			X: []float64{1.0, 2.0},
			// [R] write.table(format(shiftMCpos(MoleculeCoords, c(1.0, 2.0)), digits=22, trim=T),
			//       file="ShiftMCpos_01.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
			expected: LoadFromCsvFile2Dim("./test_data/ShiftMCpos_01.csv", ','),
		},
	}

	mc := LoadMoleculeCoordinates("./data")
	for _, tc := range testCases {
		actual := ShiftMCpos(mc, tc.X)
		if len(actual) != len(tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
		for i := 0; i < len(actual); i++ {
			for j := 0; j < len(actual[0]); j++ {
				if math.Abs(actual[i][j]-tc.expected[i][j]) > 1.0E-10 {
					t.Errorf("\n[%d][%d]\ngot  %.22f\nwant %.22f", i, j, actual[i][j], tc.expected[i][j])
					return
				}
			}
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
