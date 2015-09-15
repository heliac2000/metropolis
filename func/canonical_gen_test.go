package functions_test

import (
	"testing"

	. "./"
)

func TestCanonicalGen(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json")

	canonicalOut, characterOut, orientationOut := CanonicalGen()
	t.Logf("\nCanonical: %v\nCharacter: %v\nOrientation: %v\n",
		canonicalOut, characterOut, orientationOut)

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
