//
// prepare_init_data.go
//

package functions

import . "../util"

func PrepareInitData(ucFile, lvFile string) {
	UnitCell := LoadFromCsvFile2Dim(ucFile, ' ')
	LatticeVectors := LoadFromCsvFile2Dim(lvFile, ',')

	Lattice, Character := LatticeGen(UnitCell, LatticeVectors)

	for i := 0; i < len(Character); i++ {
		if Character[i] >= 1 && Character[i] <= 3 {
			Character[i] = 8 - Character[i]
		}
	}

	// Identify the unit cells by those which have character == 4
	wh4 := make([]int, 0, len(Character))
	for i := 0; i < len(Character); i++ {
		if Character[i] == 4 {
			wh4 = append(wh4, i)
		}
	}

	// Label the unit cell points
	nUC := len(wh4)
	UnitCellCoords := Create2DimArray(float64(0), nUC, 2).([][]float64)
	for k := 0; k < nUC; k++ {
		copy(UnitCellCoords[k], Lattice[wh4[k]])
	}

	//Adj := Create2DimArray(nUC, nUC)
	avec, bvec := LatticeVectors[0], LatticeVectors[1]
	Moves := Create2DimArray(float64(0), 4, 2).([][]float64)
	for j := 0; j < nUC; j++ {
		for i := 0; i < 4; i++ {
			copy(Moves[i], UnitCellCoords[j])
		}
		Moves[0][0] += avec[0]
		Moves[1][0] -= avec[0]
		Moves[2][0] += bvec[0]
		Moves[2][1] += bvec[1]
		Moves[3][0] -= bvec[0]
		Moves[3][1] -= bvec[1]
	}

	return
}
