package functions_test

import (
	"math"
	"testing"

	. "../util"
	. "./"
)

type testCasesCoordsIsland struct {
	CoutX, CoutC, CoutO []int
	expected            [][]float64
}

// [R] write.table(format(UnitCell2, digits=22, trim=T), file="UnitCell2.csv",
//                 sep=",", row.names=FALSE, col.names=FALSE, quote=F)
var unitCell2 [][]float64 = LoadFromCsvFile2Dim("./data/UnitCell2.csv", ',')
var unitCellCoords [][]float64 = LoadFromCsvFile2Dim("./data/UnitCellCoords.csv", ',')
var mc *MoleculeCoordinates = LoadMoleculeCoordinates("./data/Ccarts", "./data/Hcarts", "./data/Brcarts")

func TestCoordsIsland(t *testing.T) {
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
			CoutX:    []int{2},
			CoutC:    []int{6},
			CoutO:    []int{150},
			expected: LoadFromCsvFile2Dim("./data/CoordsIsland_01.csv", ','),
		},
	}

	for _, tc := range testCases {
		actual := CoordsIsland(tc.CoutX, tc.CoutC, tc.CoutO, unitCell2, unitCellCoords, mc)
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
