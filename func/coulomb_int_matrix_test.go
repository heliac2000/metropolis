package functions_test

import (
	"math"
	"path"
	"testing"

	. "../util"
	. "./"
)

type testCasesCoulombIntMatrix struct {
	k1, k2, ch1, ch2 int
	o1, o2           float64
	cm               [][]float64
	flag             bool
}

func TestCoulombIntMatrix(t *testing.T) {
	dataDir := "./data"
	SetInitData(dataDir)

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ch1 and ch2 are characters, 1-base in R but 0-base in Golang
	//
	// [R]
	// write.table(format(CoulombIntMatrix(31, 63, 1, 1, 2, 1)[[1]], digits=22, trim=T),
	//             file="CoulombIntMatrix_01.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//
	testCases := []testCasesCoulombIntMatrix{
		{
			k1: 31, k2: 63, ch1: 0, ch2: 0, o1: 2, o2: 1,
			cm:   LoadFromCsvFile2Dim(path.Join(dataDir, "CoulombIntMatrix_01.csv"), ','),
			flag: true,
		},
	}

	for _, tc := range testCases {
		actual, flag := CoulombIntMatrix(tc.k1, tc.k2, tc.ch1, tc.ch2, tc.o1, tc.o2)
		if flag != tc.flag {
			t.Errorf("got %v\nwant %v", flag, tc.flag)
			return
		}
		for i := 0; i < len(actual); i++ {
			for j := 0; j < len(actual[0]); j++ {
				if math.Abs(actual[i][j]-tc.cm[i][j]) > 1.0E-10 {
					t.Errorf("\n[%d][%d]\ngot  %.22f\nwant %.22f", i, j, actual[i][j], tc.cm[i][j])
					return
				}
			}
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
