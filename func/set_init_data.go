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
func SetInitData(ucFile, uc2File, lvFile, krlsLogFile, krlsAttFile, svmModelFile string) {
	unitCell := LoadFromCsvFile2Dim(ucFile, ' ')
	unitCell2 := LoadFromCsvFile2Dim(uc2File, ',')
	LatticeVectors := LoadFromCsvFile2Dim(lvFile, ',')

	Lattice, character := LatticeGen(unitCell, LatticeVectors)

	// Make identical unit cell points?
	// 1 <-> 11, 4 <-> 8
	for i := 0; i < len(character); i++ {
		if character[i] == 0 {
			character[i] = 10
		} else if character[i] == 3 {
			character[i] = 7
		}
	}

	// Identify the unit cells by those which have character == 4
	whC := make([]int, 0, len(character))
	for i := 0; i < len(character); i++ {
		if character[i] == CentralPoint {
			whC = append(whC, i)
		}
	}

	// Label the unit cell points
	// Make adjacency matrix for the unit cells

	// Number of unit cells
	nUC := len(whC)
	unitCellCoords := Create2DimArrayFloat(nUC, 2)
	for k := 0; k < nUC; k++ {
		copy(unitCellCoords[k], Lattice[whC[k]])
	}

	Moves, Adj := Create2DimArrayFloat(4, 2), Create2DimArrayInt(nUC, nUC)
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
	chUnique := Unique(character)
	charactersOrientations := Create2DimArrayInt(len(chUnique)*3, 3)
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

	// Load SVM objects
	LoadDataFromJSONFile(&SvmModel, svmModelFile)
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
	Zcoulomb = Create2DimArrayFloat(totAtoms, totAtoms)
	for k := 0; k < totAtoms; k++ {
		for j := 0; j < totAtoms; j++ {
			Zcoulomb[k][j] = atNum[k] * atNum[j]
			if k == j {
				Zcoulomb[k][j] = 0.5 * math.Pow(atNum[k], 2.4)
			}
		}
	}
}
