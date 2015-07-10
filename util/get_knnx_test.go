package util_test

import (
	"math"
	"reflect"
	"testing"

	. "./"
)

type testCasesGetKnnx struct {
	data, query [][]float64
	k           int
	dist        [][]float64
	index       [][]int
}

func TestGetKnnx(t *testing.T) {
	testCases := []testCasesGetKnnx{
		{
			//
			// [R] write.table(format(UnitCellCoords, digits=22, trim=T),
			//                 file="UnitCellCoords.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
			//     write.table(format(Moves, digits=22, trim=T),
			//                 file="Moves.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
			//
			data:  LoadFromCsvFile2Dim("./data/UnitCellCoords.csv", ','),
			query: LoadFromCsvFile2Dim("./data/Moves.csv", ','),
			k:     1,
			index: [][]int{{625}, {600}, {600}, {624}},
			dist: [][]float64{
				{2.518109000000002595243e+00}, {3.552713678800500929356e-15},
				{2.519199334913579502171e+00}, {7.105427357601001858711e-15},
			},
		},
	}

	for _, l := range testCases {
		dist, index := GetKnnx(l.data, l.query, l.k)

		if !reflect.DeepEqual(index, l.index) {
			t.Errorf("\ngot  %v\nwant %v", index, l.index)
			return
		}

		for i := 0; i < len(dist); i++ {
			for j := 0; j < len(dist[0]); j++ {
				if math.Abs(dist[i][j]-l.dist[i][j]) > 1.0E-14 {
					t.Errorf("\ngot  [%d][%d] %v\nwant %v", i, j, dist[i][j], l.dist[i][j])
					return
				}
			}
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
