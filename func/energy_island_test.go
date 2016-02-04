package functions_test

import (
	"testing"

	. "./"
)

type testCasesEnegyIsland struct {
	pos, chr []int
	ori      []float64
	eInt     float64
}

func TestEnergyIsland(t *testing.T) {
	SetInitData("./data")

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
	testCases := []testCasesEnegyIsland{
		{
			pos:  []int{313},
			chr:  []int{3},
			ori:  []float64{90},
			eInt: -1.146029999999996107363,
		},
		{
			pos:  []int{313},
			chr:  []int{3},
			ori:  []float64{0},
			eInt: -1.163800000000037471182,
		},
		{
			pos:  []int{313},
			chr:  []int{5},
			ori:  []float64{60},
			eInt: -1.161839999999983774615,
		},
	}

	for _, tc := range testCases {
		eInt := EnergyIsland(tc.pos, tc.chr, tc.ori)
		//if math.Abs(tc.eInt-eInt) > 1.0E-12 {
		if tc.eInt != eInt {
			t.Errorf("\nExpected Energy = %.22f\n  Actual Energy = %.22f\n", tc.eInt, eInt)
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
