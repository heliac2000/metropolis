package functions_test

import (
	"math"
	"testing"

	. "./"
)

type testCasesEnergyPairReduce struct {
	k1, k2, c1, c2 int
	o1, o2         float64
	expected       float64
}

func TestEnergyPairReduce(t *testing.T) {
	SetInitData("./data")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ch1 and ch2 are characters, 1-base in R but 0-base in Golang
	//
	// [R] EnergyPairReduce(1, 2, 3, 4, 5, 6)
	//
	testCases := []testCasesEnergyPairReduce{
		{
			k1: 1, k2: 2, c1: 2, c2: 3, o1: 5, o2: 6,
			expected: 10.0,
		},
	}

	for _, tc := range testCases {
		actual := EnergyPairReduce(tc.k1, tc.k2, tc.c1, tc.c2, tc.o1, tc.o2)
		if math.Abs(actual-tc.expected) > 1.0E-10 {
			t.Errorf("\ngot %.22f\nwant %.22f", actual, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
