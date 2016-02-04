package functions_test

import (
	"testing"

	. "./"
)

type testCasesEnegyIslandPCA struct {
	pos, chr []int
	ori      []float64
	eInt     float64
}

func TestEnergyIslandPCA(t *testing.T) {
	SetInitData("./data")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   chr is characters, 1-base in R but 0-base in Golang
	//
	// [R]
	// c <- Canonical.Order(Canonical.Gen())
	// island <- list()
	// island[[1]] = c[[1]][[1]] ## 1227 # position
	// island[[2]] = c[[2]][[1]] ## 6    # character
	// island[[3]] = c[[3]][[1]] ## 120  # orientation
	// format(EnergyIslandPCA(island), digit=22)
	// [1] "-2.037550000000123873178"
	//
	testCases := []testCasesEnegyIslandPCA{
		{
			pos:  []int{1227},
			chr:  []int{5},
			ori:  []float64{120},
			eInt: -2.037550000000123873178,
		},
		{
			pos:  []int{1227, 1227},
			chr:  []int{5, 1},
			ori:  []float64{120, 60},
			eInt: 5.935699999999769715942,
		},
		{
			pos:  []int{1227, 1227, 1227},
			chr:  []int{5, 1, 8},
			ori:  []float64{120, 60, 120},
			eInt: 23.91498999999964780727,
		},
	}

	for _, tc := range testCases {
		eInt := EnergyIslandPCA(tc.pos, tc.chr, tc.ori)
		if tc.eInt != eInt {
			t.Errorf("\nExpected Energy = %.22f\n  Actual Energy = %.22f\n", tc.eInt, eInt)
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
