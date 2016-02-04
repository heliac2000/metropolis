package functions_test

import (
	"math"
	"path"
	"testing"

	. "../util"
	. "./"
)

type testCasesCoulombVectorise struct {
	cmTest   [][]float64
	expected []float64
}

func TestCoulombVectorise(t *testing.T) {
	dataDir := "./data"
	SetInitData(dataDir)

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ch1 and ch2 are characters, 1-base in R but 0-base in Golang
	//
	// [R]
	// write.table(format(CoulombVectorise(CoulombIntMatrix(500,400,1,1,90,90)[[1]]), digits=22, trim=T),
	//             file="CoulombVectorise_01.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//
	testCases := []testCasesCoulombVectorise{
		{
			cmTest:   func() [][]float64 { ret, _ := CoulombIntMatrix(500, 400, 0, 0, 90, 90); return ret }(),
			expected: LoadFromCsvFile(path.Join(dataDir, "CoulombVectorise_01.csv"), ','),
		},
	}

	for _, tc := range testCases {
		actual := CoulombVectorise(tc.cmTest)
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
