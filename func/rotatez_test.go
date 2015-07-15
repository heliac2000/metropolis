package functions_test

import (
	"math"
	"testing"

	. "../util"
	. "./"
)

type testCasesRotateZ struct {
	theta    float64
	expected [][]float64
}

func TestRotateZ(t *testing.T) {
	testCases := []testCasesRotateZ{
		{
			theta: 0.25,
			// [R] write.table(format(RotateZ(MoleculeCoords, 0.25), digits=22, trim=T),
			//     file="RotateZ_01.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
			expected: LoadFromCsvFile2Dim("./data/RotateZ_01.csv", ','),
		},
	}

	mc := LoadMoleculeCoordinates("./data/Ccarts", "./data/Hcarts", "./data/Brcarts")
	for _, tc := range testCases {
		actual := RotateZ(mc, tc.theta)
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
