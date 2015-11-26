package functions_test

import (
	"testing"

	. "./"
)

type testCasesEnegyCanonical struct {
	pos, chr [][]int
	ori      [][]float64
	eInt     float64
}

func TestEnergyCanonical(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   chr is characters, 1-base in R but 0-base in Golang
	//
	// [R]
	// c <- Canonical.Order(Canonical.Gen())
	// island <- list()
	// island[[1]] = c[[1]][[1]] ## 313 # position
	// island[[2]] = c[[2]][[1]] ## 4   # character
	// island[[3]] = c[[3]][[1]] ## 90  # orientation
	// format(EnergyIsland(island), digit=22)
	// [1] "-1.146029999999996107363"
	//
	testCases := []testCasesEnegyCanonical{
		{
			pos:  [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			chr:  [][]int{{3}, {3}, {5}, {5}, {4}, {4}, {4}, {5}, {3}, {4}, {0}},
			ori:  [][]float64{{90}, {0}, {60}, {120}, {30}, {120}, {60}, {0}, {90}, {30}, {0}},
			eInt: -11.52151000000014846592,
		},
	}

	for _, tc := range testCases {
		eInt := EnergyCanonical(tc.pos, tc.chr, tc.ori)
		//if math.Abs(tc.eInt-eInt) > 1.0E-12 {
		if tc.eInt != eInt {
			t.Errorf("\nExpected Energy = %.22f\n  Actual Energy = %.22f\n", tc.eInt, eInt)
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
