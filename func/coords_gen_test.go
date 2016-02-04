package functions_test

import (
	"math"
	"path"
	"testing"

	. "../util"
	. "./"
)

type testCasesCoordsGen struct {
	k1, k2, ch1, ch2 int
	o1, o2           float64
	expected         [][]float64
}

func TestCoordsGen(t *testing.T) {
	dataDir := "./data"
	SetInitData(dataDir)

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ch1, ch2: 1-base in R but 0-base in Golang
	//
	// [R]
	// write.table(format(cGen(313, 363, 6, 6, 0, 0), digits=22, trim=T),
	//             file="CoordsGen_01.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//
	testCases := []testCasesCoordsGen{
		{
			k1: 313, k2: 363, ch1: 5, ch2: 5, o1: 0, o2: 0,
			expected: LoadFromCsvFile2Dim(path.Join(dataDir, "CoordsGen_01.csv"), ','),
		},
	}

	for _, tc := range testCases {
		actual := CoordsGen(tc.k1, tc.k2, tc.ch1, tc.ch2, tc.o1, tc.o2)
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

// [R]
// cGen <- function(k1, k2, c1, c2, o1, o2) # Turn pair into molecule coordinates
// 	{
// 	 CoutX <- c(k1, k2)
// 	 CoutC <- c(c1, c2)
// 	 CoutO <- c(o1, o2)
//
//   IslandTemp <- array(0, c(0,3))
// 	for(h in 1:length(CoutX))
// 		{deltaxy1 <- c(UnitCell[central.point,2] - UnitCell[CoutC[h],2], UnitCell[central.point,3] - UnitCell[CoutC[h],3])
// 		coords1 = UnitCellCoords[CoutX[h],] - deltaxy1
// 		M1 = shiftMCpos(MoleculeCoords,coords1)
// 	 	IslandTemp <- rbind(IslandTemp, RotateZ(M1,CoutO[h]*pi/180))
// 		}
// 	IslandTemp
// 	}
