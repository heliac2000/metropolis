package functions_test

import (
	"testing"

	. "./"
)

func TestRandomIslandUnitCell(t *testing.T) {
	unitCell, _, adjCuml, chUnique :=
		PrepareInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	islandP, islandC, islandO := RandomIslandUnitCell(4, adjCuml[Npower-1], unitCell, chUnique)
	t.Logf("\nIslandP: %v\nIslandC: %v\nIslandO: %v\n", islandP, islandC, islandO)

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
