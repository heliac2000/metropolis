//
// prepare_init_data.go
//

package functions

import . "../util"

func PrepareInitData(ucFile, lvFile string) ([][]float64, [][]float64, [][][]int, []int) {
	UnitCell := LoadFromCsvFile2Dim(ucFile, ' ')
	LatticeVectors := LoadFromCsvFile2Dim(lvFile, ',')

	Lattice, Character := LatticeGen(UnitCell, LatticeVectors)

	// Make identical unit cell points?
	// 1 <-> 11, 4 <-> 8
	for i := 0; i < len(Character); i++ {
		if Character[i] == 0 {
			Character[i] = 10
		} else if Character[i] == 3 {
			Character[i] = 7
		}
	}

	// Identify the unit cells by those which have character == central.point
	whC := make([]int, 0, len(Character))
	for i := 0; i < len(Character); i++ {
		if Character[i] == CentralPoint {
			whC = append(whC, i)
		}
	}

	// Label the unit cell points
	// Make adjacency matrix for the unit cells

	// Number of unit cells
	nUC := len(whC)
	UnitCellCoords := Create2DimArrayFloat(nUC, 2)
	for k := 0; k < nUC; k++ {
		copy(UnitCellCoords[k], Lattice[whC[k]])
	}

	Moves, Adj := Create2DimArrayFloat(4, 2), Create2DimArrayInt(nUC, nUC)
	avec, bvec := LatticeVectors[0], LatticeVectors[1]
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
	AdjCuml[0] = Copy2DimArrayInt(AdjSEQ[0])
	for k := 1; k < Npower; k++ {
		AdjCuml[k] = Copy2DimArrayInt(AdjSEQ[0])
		for j := 1; j <= k; j++ {
			MatrixAdd(AdjCuml[k], AdjSEQ[j])
		}
	}

	return UnitCell, UnitCellCoords, AdjCuml, Unique(Character)
}
