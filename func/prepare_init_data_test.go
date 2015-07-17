package functions_test

import (
	"math"
	"reflect"
	"testing"

	. "../util"
	. "./"
)

func TestPrepareInitData(t *testing.T) {
	adjCuml, unitCellCoords := PrepareInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")
	// [R] writeListData(AdjCuml, "AdjCuml.csv")
	adjCumlR := LoadFromCsvFileList("./data/AdjCuml.csv")
	// [R] write.table(format(UnitCellCoords, digits=22, trim=T),
	//                 file="UnitCellCoords.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	unitCellCoordsR := LoadFromCsvFile2Dim("./data/UnitCellCoords.csv", ',')

	if !reflect.DeepEqual(adjCuml, adjCumlR) {
		t.Errorf("\ngot  %v\nwant %v", adjCuml, adjCumlR)
		return
	}

	if len(unitCellCoords) != len(unitCellCoordsR) {
		t.Errorf("\ngot  %v\nwant %v", unitCellCoords, unitCellCoordsR)
		return
	}
	for i := 0; i < len(unitCellCoords); i++ {
		for j := 0; j < len(unitCellCoords[0]); j++ {
			if math.Abs(unitCellCoords[i][j]-unitCellCoordsR[i][j]) > 1.0E-10 {
				t.Errorf("\n[%d][%d]\ngot  %.22f\nwant %.22f", i, j, unitCellCoords[i][j], unitCellCoordsR[i][j])
				return
			}
		}
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
