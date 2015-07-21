//
// set_init_data.go
//

package functions

import . "../util"

type InitData struct {
	UnitCell       [][]float64
	UnitCellCoords [][]float64
	AdjCuml        [][][]int
	Character      []int
	ChUnique       []int
}

func SetInitData(ucFile, lvFile string) *InitData {
	UnitCell := LoadFromCsvFile2Dim(ucFile, ' ')
	LatticeVectors := LoadFromCsvFile2Dim(lvFile, ',')

	Lattice, Character := LatticeGen(UnitCell, LatticeVectors)

	for i := 0; i < len(Character); i++ {
		if Character[i] >= 0 && Character[i] <= 2 {
			Character[i] = 6 - Character[i]
		}
	}

	// Identify the unit cells by those which have character == 4
	wh4 := make([]int, 0, len(Character))
	for i := 0; i < len(Character); i++ {
		if Character[i] == 3 {
			wh4 = append(wh4, i)
		}
	}

	// Label the unit cell points
	// Make adjacency matrix for the unit cells

	// Number of unit cells
	nUC := len(wh4)
	UnitCellCoords := Create2DimArray(float64(0), nUC, 2).([][]float64)
	for k := 0; k < nUC; k++ {
		copy(UnitCellCoords[k], Lattice[wh4[k]])
	}

	avec, bvec := LatticeVectors[0], LatticeVectors[1]
	Moves := Create2DimArray(float64(0), 4, 2).([][]float64)
	Adj := Create2DimArray(int(0), nUC, nUC).([][]int)
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

		_, surrj := GetKnnx(UnitCellCoords, Moves, 1)
		surr := make([]int, len(surrj))
		for i := 0; i < len(surrj); i++ {
			surr[i] = surrj[i][0]
		}
		surr = Unique(surr)

		for i := 0; i < len(surr); i++ {
			Adj[j][surr[i]] = 1
		}
		Adj[j][j] = 0
	}

	// Make a sequence of power matrices
	AdjSEQ := make([][][]int, Npower)
	AdjSEQ[0] = Adj
	for k := 1; k < Npower; k++ {
		// Generate the power matrices
		AdjSEQ[k] = MatrixMultiply(AdjSEQ[k-1], Adj)
	}

	AdjCuml := make([][][]int, Npower)
	AdjCuml[0] = Copy2DimArray(AdjSEQ[0]).([][]int)
	for k := 1; k < Npower; k++ {
		AdjCuml[k] = Copy2DimArray(AdjSEQ[0]).([][]int)
		for j := 1; j <= k; j++ {
			MatrixAdd(AdjCuml[k], AdjSEQ[j])
		}
	}

	return &InitData{
		UnitCell:       UnitCell,
		UnitCellCoords: UnitCellCoords,
		AdjCuml:        AdjCuml,
		Character:      Character,
		ChUnique:       Unique(Character),
	}
}
