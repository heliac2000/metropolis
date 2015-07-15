package functions_test

import (
	"math"
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testCasesLatticeGen struct {
	UC, LV    [][]float64
	Lattice   [][]float64
	Character []int
}

func TestLatticeGen(t *testing.T) {
	testCases := []testCasesLatticeGen{
		{
			UC: LoadFromCsvFile2Dim("./data/PrecursorUnitCell.csv", ' '),
			LV: LoadFromCsvFile2Dim("./data/PrecursorUnitCellAxes.csv", ','),
			//
			// [R] write.table(Lattice, file="lattice.csv", sep=",", row.names=FALSE, col.names=FALSE)
			Lattice: LoadFromCsvFile2Dim("./data/lattice.csv", ','),
			//
			// [R] cc = LatticeGen(UnitCell, LatticeVectors)[[2]]
			//     write.table(Lattice, file="character.dat", row.names=FALSE, col.names=FALSE)
			Character: LoadFromCsvFileInt("./data/character.dat"),
		},
	}

	for _, l := range testCases {
		lat, char := LatticeGen(l.UC, l.LV)
		if len(lat) != len(l.Lattice) {
			t.Errorf("\ngot  %v, %v\nwant %v, %v", lat, char, l.Lattice, l.Character)
			return
		}
		for i := 0; i < len(lat); i++ {
			for j := 0; j < len(lat[0]); j++ {
				if math.Abs(lat[i][j]-l.Lattice[i][j]) > 1.0E-10 {
					t.Errorf("\ngot  [%d][%d] %v\nwant %v", i, j, lat[i][j], l.Lattice[i][j])
					return
				}
			}
		}

		if !reflect.DeepEqual(char, l.Character) {
			t.Errorf("\ngot  %v\nwant %v", char, l.Character)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
