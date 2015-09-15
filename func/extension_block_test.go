package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testCasesExtensionBlock struct {
	xtest       []int
	zcoords     [][]float64
	xtestAppend [][][]int
	lx          int
}

func TestMakeExtensionBlock(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json")

	testCases := []testCasesExtensionBlock{
		{
			// [R] t1 = Extension.Block(list(c(10,20,30), c(5,3,4), c(6,7,8)),
			//            array(c(1.0, 2.0, 3.0, 4.0, 5.0, 6.0), c(2, 3)))[[1]]
			//     ## R is 1-base index, golang is 0-base.
			//     for (i in 1:length(t1)){ t1[[i]][,2] = t1[[i]][,2] - 1 }
			//     writeListData(t1, "ExtensionBlock_01.csv")
			//
			xtest:       []int{10, 20, 30},
			zcoords:     [][]float64{{1.0, 3.0, 5.0}, {2.0, 4.0, 6.0}},
			xtestAppend: LoadFromCsvFileList("./data/ExtensionBlock_01.csv"),
			lx:          1176,
		},
	}

	for _, tc := range testCases {
		xtestAppend, lx := ExtensionBlock(tc.xtest, tc.zcoords)
		if !reflect.DeepEqual(xtestAppend, tc.xtestAppend) || lx != tc.lx {
			t.Errorf("\ngot  %v\nwant %v", lx, tc.lx)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
