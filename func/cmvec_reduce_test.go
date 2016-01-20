package functions_test

import (
	"math"
	"testing"

	. "./"
)

type testCasesCMvecReduce struct {
	cmVec    []float64
	expected []float64
}

func TestCoulombVectorise(t *testing.T) {
	SetInitData()

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ch1 and ch2 are characters, 1-base in R but 0-base in Golang
	//
	// [R]
	// npcs = 7
	// CMvecReduce(CoulombVectorise(CoulombIntMatrix(500,400,1,1,90,90)[[1]]), npcs)
	//
	testCases := []testCasesCMvecReduce{
		{
			cmVec: func() []float64 {
				ci, _ := CoulombIntMatrix(500, 400, 0, 0, 90, 90)
				cv := CoulombVectorise(ci)
				return cv
			}(),
			expected: []float64{
				181.83229427688087298520, 52.07623191468651668856, -88.78906289938313989296, -272.03612236354888409551,
				-153.47368025052372786377, 152.49316005940150375864, -135.26865759958911894500,
			},
		},
	}

	for _, tc := range testCases {
		actual := CMvecReduce(tc.cmVec, Npcs)
		for i := 0; i < len(actual); i++ {
			if math.Abs(actual[i]-tc.expected[i]) > 1.0E-10 {
				t.Errorf("\n[%d]\ngot  %.22f\nwant %.22f", i, actual[i], tc.expected[i])
				return
			}
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
