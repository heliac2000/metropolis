//
// set_init_data.go
//

package functions

import . "../util"

type InitData struct {
	UnitCell               [][]float64
	UnitCellCoords         [][]float64
	AdjCuml                [][][]int
	Character              []int
	ChUnique               []int
	CharactersOrientations [][]int
}

func SetInitData(ucFile, lvFile string) {
	unitCell := LoadFromCsvFile2Dim(ucFile, ' ')
	LatticeVectors := LoadFromCsvFile2Dim(lvFile, ',')

	Lattice, character := LatticeGen(unitCell, LatticeVectors)

	for i := 0; i < len(character); i++ {
		if character[i] >= 0 && character[i] <= 2 {
			character[i] = 6 - character[i]
		}
	}

	// Identify the unit cells by those which have character == 4
	wh4 := make([]int, 0, len(character))
	for i := 0; i < len(character); i++ {
		if character[i] == 3 {
			wh4 = append(wh4, i)
		}
	}

	// Label the unit cell points
	// Make adjacency matrix for the unit cells

	// Number of unit cells
	nUC := len(wh4)
	unitCellCoords := Create2DimArray(float64(0), nUC, 2).([][]float64)
	for k := 0; k < nUC; k++ {
		copy(unitCellCoords[k], Lattice[wh4[k]])
	}

	avec, bvec := LatticeVectors[0], LatticeVectors[1]
	Moves := Create2DimArray(float64(0), 4, 2).([][]float64)
	Adj := Create2DimArray(int(0), nUC, nUC).([][]int)
	for j := 0; j < nUC; j++ {
		for i := 0; i < 4; i++ {
			copy(Moves[i], unitCellCoords[j])
		}
		Moves[0][0] += avec[0]
		Moves[1][0] -= avec[0]
		Moves[2][0] += bvec[0]
		Moves[2][1] += bvec[1]
		Moves[3][0] -= bvec[0]
		Moves[3][1] -= bvec[1]

		_, surrj := GetKnnx(unitCellCoords, Moves, 1)
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

	adjCuml := make([][][]int, Npower)
	adjCuml[0] = Copy2DimArray(AdjSEQ[0]).([][]int)
	for k := 1; k < Npower; k++ {
		adjCuml[k] = Copy2DimArray(AdjSEQ[0]).([][]int)
		for j := 1; j <= k; j++ {
			MatrixAdd(adjCuml[k], AdjSEQ[j])
		}
	}

	// Generate all combinations of characters and corresponding
	// orientations. First space is empty - it gets filled in
	// Extension.Block
	//
	// charactersOrientations[][1] is index(0-base or 1-base)
	chUnique := Unique(character)
	charactersOrientations := Create2DimArray(int(0), len(chUnique)*3, 3).([][]int)
	for k, cnt := 0, 0; k < len(chUnique); k++ {
		for j := 0; j < 3; j++ {
			charactersOrientations[cnt][1] = chUnique[k]
			charactersOrientations[cnt][2] = int(unitCell[chUnique[k]][j+4])
			cnt++
		}
	}

	Inp = &InitData{
		UnitCell:               unitCell,
		UnitCellCoords:         unitCellCoords,
		AdjCuml:                adjCuml,
		Character:              character,
		ChUnique:               chUnique,
		CharactersOrientations: charactersOrientations,
	}
}
