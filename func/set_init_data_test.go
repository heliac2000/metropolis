package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

//
// [R] write.table(format(Zcoulomb, digits=22, trim=T),
//       file="Zcoulomb.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
//
func TestSetZcoulomb(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv")
	zcb := LoadFromCsvFile2Dim("./data/Zcoulomb.csv", ',')

	if !reflect.DeepEqual(Zcoulomb, zcb) {
		t.Errorf("\ngot  %v\nwant %v", Zcoulomb, zcb)
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
