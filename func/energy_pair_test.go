package functions_test

import (
	"math"
	"testing"

	. "./"
)

type testCasesEnegyPair struct {
	k1, k2, ch1, ch2 int
	o1, o2           float64
	intType          string
	eint             float64
}

func TestEnergyPair(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// [R]
	//
	// format(EnergyPair(31, 63, 1, 1, 2, 1), digits=22, trim=T)
	//
	testCases := []testCasesEnegyPair{
		{
			k1: 313, k2: 363, ch1: 6, ch2: 6, o1: 0, o2: 0,
			intType: "repulsive", eint: 0.9265127673648503314752,
		},
		{
			k1: 31, k2: 63, ch1: 1, ch2: 1, o1: 2, o2: 1,
			intType: "attractive", eint: -0.05134541334315906313535,
		},
	}

	for _, tc := range testCases {
		intType, eint := EnergyPair(tc.k1, tc.k2, tc.ch1, tc.ch2, tc.o1, tc.o2)
		if tc.intType != intType ||
			math.Abs(tc.eint-eint) > 10E-12 {
			t.Errorf("\nExpected: Type   = %v, Energy = %v\n  Actual: Type   = %v, Energy = %v\n",
				tc.intType, tc.eint, intType, eint)
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
