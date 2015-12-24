package functions_test

import (
	"path"
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
	SetInitData()
	zcb := LoadFromCsvFile2Dim(path.Join(DATA_DIR, "Zcoulomb.csv"), ',')

	if !reflect.DeepEqual(Zcoulomb, zcb) {
		t.Errorf("\ngot  %v\nwant %v", Zcoulomb, zcb)
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
