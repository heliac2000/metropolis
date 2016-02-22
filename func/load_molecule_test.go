package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

func TestLoadMoleculeCoordinates(t *testing.T) {
	mc := LoadMoleculeCoordinates("./data", "Molecule_01.csv", "Molecule_02.csv", "Molecule_03.csv")

	// [R]
	//  write.table(format(C, digits=22, trim=T),
	//                 file="C.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//  write.table(format(H, digits=22, trim=T),
	//                 file="H.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//  write.table(format(Br, digits=22, trim=T),
	//                 file="Br.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	expected := MoleculeCoordinates{
		C:  LoadFromCsvFile2Dim("./test_data/C.csv", ','),
		H:  LoadFromCsvFile2Dim("./test_data/H.csv", ','),
		Br: LoadFromCsvFile2Dim("./test_data/Br.csv", ','),
	}

	if !reflect.DeepEqual(mc.C, expected.C) ||
		!reflect.DeepEqual(mc.H, expected.H) ||
		!reflect.DeepEqual(mc.Br, expected.Br) {
		t.Errorf("\ngot  %v\nwant %v", *mc, expected)
		return
	}

}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
