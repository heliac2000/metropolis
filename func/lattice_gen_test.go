package functions_test

import (
	"math"
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testCasesLatticeGen struct {
	uc, lv    [][]float64
	lattice   [][]float64
	character []int
}

func TestLatticeGen(t *testing.T) {
	testCases := []testCasesLatticeGen{
		{
			uc: LoadFromCsvFile2Dim("./data/PrecursorUnitCell.csv", ' '),
			lv: LoadFromCsvFile2Dim("./data/PrecursorUnitCellAxes.csv", ','),
			//
			// [R] write.table(lattice, file="lattice.csv", sep=",", row.names=FALSE, col.names=FALSE)
			lattice: LoadFromCsvFile2Dim("./data/lattice.csv", ','),
			//
			// NOTICE: R is 1-base index, golang is 0-base.
			//
			// [R] cc = LatticeGen(UnitCell, latticeVectors)[[2]] - 1
			//     write.table(cc, file="character.dat", row.names=FALSE, col.names=FALSE)
			character: LoadFromCsvFileInt("./data/character.dat"),
		},
	}

	for _, l := range testCases {
		lat, char := LatticeGen(l.uc, l.lv)
		if len(lat) != len(l.lattice) {
			t.Errorf("\ngot  %v, %v\nwant %v, %v", lat, char, l.lattice, l.character)
			return
		}
		for i := 0; i < len(lat); i++ {
			for j := 0; j < len(lat[0]); j++ {
				if math.Abs(lat[i][j]-l.lattice[i][j]) > 1.0E-10 {
					t.Errorf("\ngot  [%d][%d] %v\nwant %v", i, j, lat[i][j], l.lattice[i][j])
					return
				}
			}
		}

		if !reflect.DeepEqual(char, l.character) {
			t.Errorf("\ngot  %v\nwant %v", char, l.character)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
