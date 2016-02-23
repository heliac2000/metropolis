package functions_test

import (
	"math"
	"testing"

	. "../util"
	. "./"
)

//
// [R] write.table(format(Zcoulomb, digits=22, trim=T),
//       file="Zcoulomb.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
//
func TestSetZcoulomb(t *testing.T) {
	SetInitData("./data")
	zcb := LoadFromCsvFile2Dim("test_data/Zcoulomb.csv", ',')

	for i := 0; i < len(zcb); i++ {
		for j := 0; j < len(zcb[i]); j++ {
			if math.Abs(zcb[i][j]-Zcoulomb[i][j]) > 1.0E-12 {
				t.Errorf("got  %.22f\nwant %.22f", zcb[i][j], Zcoulomb[i][j])
				return
			}
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
