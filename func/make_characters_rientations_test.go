package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testCasesMakeCharactersRientations struct {
	zcoords  [][]float64
	xadd     []int
	expected [][]int
}

func TestMakeCharactersRientations(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	testCases := []testCasesMakeCharactersRientations{
		{
			// [R] t1 = makeCharactersOrientations(1, t(array(c(22,33,44,55,66,77),3,2)), t(as.matrix(c(1,2,3))))
			//     t1[,2] = t1[,2] - 1 ## R is 1-base index, golang is 0-base.
			//     write.table(t1, file="makeCharactersOrientations_01.dat", row.names=FALSE, col.names=FALSE, sep=",")
			//
			zcoords:  [][]float64{{22, 33, 44}, {55, 66, 77}},
			xadd:     []int{1, 2, 3},
			expected: LoadFromCsvFile2DimInt("./data/makeCharactersOrientations_01.dat", ','),
		},
	}

	for _, tc := range testCases {
		actual := MakeCharactersRientations(tc.zcoords, tc.xadd)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
