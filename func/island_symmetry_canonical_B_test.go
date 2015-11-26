package functions_test

import (
	"testing"

	. "./"
)

type testCasesISCB struct {
	pos  [][]int
	fact float64
}

func TestEnergyISCB(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// [R]
	// c <- Canonical.Order(Canonical.Gen())
	// island.symmetry.canonical.B(c)
	// [1] 1
	//
	testCases := []testCasesISCB{
		{
			pos:  [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			fact: 1.0,
		},
	}

	for _, tc := range testCases {
		fact := IslandSymmetryCanonicalB(tc.pos)
		if tc.fact != fact {
			t.Errorf("\nExpected Energy = %.22f\n  Actual Energy = %.22f\n", tc.fact, fact)
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
