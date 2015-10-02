package functions_test

import (
	"testing"

	. "./"
)

type testEnegyPair struct {
	k1, k2, ch1, ch2 int
	o1, o2           float64
}

func TestEnergyPair(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	testCases := []testEnegyPair{
		{313, 363, 6, 6, 0, 0},
		{31, 63, 1, 1, 2, 1},
	}

	for _, tc := range testCases {
		intType, eint := EnergyPair(tc.k1, tc.k2, tc.ch1, tc.ch2, tc.o1, tc.o2)
		t.Logf("\nType   = %v\nEnergy = %v\n", intType, eint)
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
