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
	SetInitData("./data")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   chr is characters, 1-base in R but 0-base in Golang
	//
	// [R]
	// c <- Canonical.Order(Canonical.Gen())
	// format(EnergyCanonical(c), digit=22)
	// [1] "-20.20290000000102281774"
	//
	testCases := []testCasesEnegyCanonical{
		{
			pos:  [][]int{{1227}, {1227}, {1227}, {1227}, {1227}, {1227}, {1227}, {1227}, {1227}, {1227}, {0}},
			chr:  [][]int{{1}, {8}, {7}, {7}, {8}, {10}, {2}, {4}, {4}, {10}, {0}},
			ori:  [][]float64{{60}, {120}, {60}, {60}, {120}, {60}, {120}, {0}, {0}, {120}, {0}},
			eInt: -20.20290000000102281774,
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
