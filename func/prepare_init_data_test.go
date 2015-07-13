package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

func TestPrepareInitData(t *testing.T) {
	// [R] writeListData(AdjCuml, "AdjCuml.csv")
	adjCuml := PrepareInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")
	adjCumlR := LoadFromCsvFileList("./data/AdjCuml.csv")
	if !reflect.DeepEqual(adjCuml, adjCumlR) {
		t.Errorf("\ngot  %v\nwant %v", adjCuml, adjCumlR)
		return
	}

	// [R] write.table(Adj, file="Adj.csv", sep=",", row.names=FALSE, col.names=FALSE)
	// adj := PrepareInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")
	// adjR := LoadFromCsvFile2DimInt("./data/Adj.csv", ',')
	// if !reflect.DeepEqual(adj, adjR) {
	// 	t.Errorf("\ngot  %v\nwant %v", adj, adjR)
	// 	return
	// }

	// [R] writeListData(AdjSEQ, "AdjSEQ.csv")
	// adjSeq := PrepareInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")
	// adjSeqR := LoadFromCsvFileList("./data/AdjSEQ.csv")
	// if !reflect.DeepEqual(adjSeq, adjSeqR) {
	// 	t.Errorf("\ngot  %v\nwant %v", adjSeq, adjSeqR)
	// 	return
	// }
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
