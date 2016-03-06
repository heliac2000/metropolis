package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

func TestLoadMoleculeCoordinates(t *testing.T) {
	mc := LoadMoleculeCoordinates("./data")

	// [R]
	//  write.table(format(C, digits=22, trim=T),
	//                 file="C.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//  write.table(format(H, digits=22, trim=T),
	//                 file="H.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//  write.table(format(Br, digits=22, trim=T),
	//                 file="Br.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	expected := MoleculeCoordinates{
		Mol: [][][]float64{
			LoadFromCsvFile2Dim("./test_data/C.csv", ','),
			LoadFromCsvFile2Dim("./test_data/H.csv", ','),
			LoadFromCsvFile2Dim("./test_data/Br.csv", ','),
		},
	}

	if !reflect.DeepEqual(mc.Mol[0], expected.Mol[0]) ||
		!reflect.DeepEqual(mc.Mol[1], expected.Mol[1]) ||
		!reflect.DeepEqual(mc.Mol[2], expected.Mol[2]) {
		t.Errorf("\ngot  %v\nwant %v", *mc, expected)
		return
	}

}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
