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
			uc: LoadFromCsvFile2Dim("./data/UnitCell.csv", ','),
			lv: LoadFromCsvFile2Dim("./data/PrecursorUnitCellAxes.csv", ','),
			//
			// NOTICE: R is 1-base index, golang is 0-base.
			//
			// [R]
			// cc = LatticeGen(UnitCell, LatticeVectors)
			// write.table(format(cc[[1]], digits=22, trim=T),
			//             file="lattice_test.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
			// write.table(cc[[2]]-1, file="character_test.dat", row.names=FALSE, col.names=FALSE, quote=F)
			lattice:   LoadFromCsvFile2Dim("./test_data/lattice_test.csv", ','),
			character: LoadFromCsvFileInt("./test_data/character_test.dat"),
		},
	}

	for _, l := range testCases {
		lat, char := LatticeGen(l.uc, l.lv)
		if len(lat) != len(l.lattice) {
			t.Errorf("\ngot  %v\nwant %v", lat, l.lattice)
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
