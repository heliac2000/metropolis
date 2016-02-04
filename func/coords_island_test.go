package functions_test

import (
	"math"
	"path"
	"testing"

	. "../util"
	. "./"
)

type testCasesCoordsIsland struct {
	coutX, coutC []int
	coutO        []float64
	expected     [][]float64
}

func TestCoordsIsland(t *testing.T) {
	dataDir := "./data"
	SetInitData(dataDir)

	testCases := []testCasesCoordsIsland{
		{
			// pt[[1]]: Positions, pt[[2]]: Characters, pt[[3]]: Orientations
			// Positions and Characters are index, Orientations is angle.
			//
			// NOTICE: R is 1-base index, golang is 0-base.
			//
			// pt <- list(); pt[[1]] = 3; pt[[2]] = 7; pt[[3]] = 150
			// [R] write.table(format(CoordsIsland(pt), digits=22, trim=T),
			//     file="CoordsIsland_01.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
			//
			coutX:    []int{2},
			coutC:    []int{6},
			coutO:    []float64{150},
			expected: LoadFromCsvFile2Dim(path.Join(dataDir, "CoordsIsland_01.csv"), ','),
		},
	}

	for _, tc := range testCases {
		actual := CoordsIsland(tc.coutX, tc.coutC, tc.coutO)
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
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
