package functions_test

import (
	"testing"

	. "./"
)

func TestRandomIslandUnitCell(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	islandP, islandC, islandO :=
		RandomIslandUnitCell(4, Inp.AdjCuml[Npower-1], Inp.UnitCell, Inp.ChUnique)
	t.Logf("\nIslandP: %v\nIslandC: %v\nIslandO: %v\n", islandP, islandC, islandO)

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
