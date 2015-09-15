//
// set_init_data.go
//

package functions

import (
	"math"

	. "../util"
)

type InitData struct {
	UnitCell               [][]float64
	UnitCell2              [][]float64
	UnitCellCoords         [][]float64
	AdjCuml                [][][]int
	Character              []int
	ChUnique               []int
	CharactersOrientations [][]int
	MoleculeCoordinates    *MoleculeCoordinates
}

//
// ucFile: UnitCell data
// uc2File: UnitCell2 data(CSV format)
// lvFile: LatticeVector data(CSV format)
//
// [R] write.table(format(UnitCell2, digits=22, trim=T), file="UnitCell2.csv",
//                 sep=",", row.names=FALSE, col.names=FALSE, quote=F)
//
func SetInitData(ucFile, uc2File, lvFile, krlsLogFile, krlsAttFile string) {
	unitCell := LoadFromCsvFile2Dim(ucFile, ' ')
	unitCell2 := LoadFromCsvFile2Dim(uc2File, ',')
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
	var unitCellCoords [][]float64
	nUC := len(wh4)
	Create2DimArray(&unitCellCoords, nUC, 2)
	for k := 0; k < nUC; k++ {
		copy(unitCellCoords[k], Lattice[wh4[k]])
	}

	var Moves [][]float64
	var Adj [][]int
	_, _ = Create2DimArray(&Moves, 4, 2), Create2DimArray(&Adj, nUC, nUC)
	avec, bvec := LatticeVectors[0], LatticeVectors[1]
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
	Copy2DimArray(&adjCuml[0], AdjSEQ[0])
	for k := 1; k < Npower; k++ {
		Copy2DimArray(&adjCuml[k], AdjSEQ[0])
		for j := 1; j <= k; j++ {
			MatrixAdd(adjCuml[k], AdjSEQ[j])
		}
	}

	// Generate all combinations of characters and corresponding
	// orientations. First space is empty - it gets filled in
	// Extension.Block
	//
	// charactersOrientations[][1] is index(0-base or 1-base)
	//
	var charactersOrientations [][]int
	chUnique := Unique(character)
	Create2DimArray(&charactersOrientations, len(chUnique)*3, 3)
	for k, cnt := 0, 0; k < len(chUnique); k++ {
		for j := 0; j < 3; j++ {
			charactersOrientations[cnt][1] = chUnique[k]
			charactersOrientations[cnt][2] = int(unitCell[chUnique[k]][j+4])
			cnt++
		}
	}

	Inp = &InitData{
		UnitCell:               unitCell,
		UnitCell2:              unitCell2,
		UnitCellCoords:         unitCellCoords,
		AdjCuml:                adjCuml,
		Character:              character,
		ChUnique:               chUnique,
		CharactersOrientations: charactersOrientations,
		MoleculeCoordinates:    LoadMoleculeCoordinates("./data/Ccarts", "./data/Hcarts", "./data/Brcarts"),
	}

	SetZcoulomb()

	// Load KRLS objects
	LoadDataFromJSONFile(&KernelRegsRepLog, krlsLogFile)
	LoadDataFromJSONFile(&KernelRegsAtt, krlsAttFile)
}

// Prepare the numerators of the Coulomb matrices
//
func SetZcoulomb() {
	totAtoms := 0
	for _, v := range Natoms {
		totAtoms += v
	}

	// Assign the atomic numbers
	atNum := make([]float64, totAtoms)
	thrNum := []int{Natoms[0], Natoms[0] + Natoms[1], totAtoms}
	for k := 0; k < totAtoms; k++ {
		for i, v := range thrNum {
			if k < v {
				atNum[k] = AtomNumber[i]
				break
			}
		}
	}

	// Prepare the numerators of the Coulomb matrices
	Create2DimArray(&Zcoulomb, totAtoms, totAtoms)
	for k := 0; k < totAtoms; k++ {
		for j := 0; j < totAtoms; j++ {
			Zcoulomb[k][j] = atNum[k] * atNum[j]
			if k == j {
				Zcoulomb[k][j] = 0.5 * math.Pow(atNum[k], 2.4)
			}
		}
	}
}
