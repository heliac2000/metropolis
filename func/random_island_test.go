package functions_test

import (
	"testing"

	. "./"
)

func TestRandomIslandUnitCell(t *testing.T) {
	SetInitData("./data")

	islandP, islandC, islandO := RandomIslandUnitCell(4)
	t.Logf("\nIslandP: %v\nIslandC: %v\nIslandO: %v\n", islandP, islandC, islandO)

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
