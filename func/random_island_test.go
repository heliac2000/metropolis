package functions_test

import (
	"fmt"
	"testing"

	. "./"
)

func TestRandomIslandUnitCell(t *testing.T) {
	unitCell, _, adjCuml, chUnique :=
		PrepareInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	islandP, islandC, islandO := RandomIslandUnitCell(2, adjCuml[Npower-1], unitCell, chUnique)

	fmt.Println(islandP)
	fmt.Println(islandC)
	fmt.Println(islandO)

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
