//
// set_init_data.go
//

package functions

import (
	"math"
	"path"
	"sort"

	. "../util"
)

type InitData struct {
	UnitCell               [][]float64
	UnitCell2              [][]float64
	OrientationsEnergies   [][][]float64
	UnitCellCoords         [][]float64
	AdjCuml                [][][]int
	Character              []int
	ChUnique               []int
	CharactersOrientations [][]int
	MoleculeCoordinates    *MoleculeCoordinates
}

//
// [R]
// write.table(format(UnitCell2, digits=22, trim=T), file="UnitCell2.csv",
//             sep=",", row.names=FALSE, col.names=FALSE, quote=F)
//
func SetInitData() {
	unitCell2 := LoadFromCsvFile2Dim(path.Join(DATA_DIR, "UnitCell2.csv"), ',')
	LatticeVectors := LoadFromCsvFile2Dim(path.Join(DATA_DIR, "PrecursorUnitCellAxes.csv"), ',')

	// Identify the colors
	c := make([]int, len(unitCell2))
	for i := 0; i < len(unitCell2); i++ {
		c[i] = int(unitCell2[i][0])
	}
	colors := Unique(c)
	sort.Ints(colors)

	// This should be ordered according to color!
	unitCell := Create2DimArrayFloat(len(colors), 3)
	for k := 0; k < len(colors); k++ {
		whk := 0
		for i, v := range c {
			if v == colors[k] {
				whk = i
				break
			}
		}
		unitCell[k][0] = float64(colors[k])
		unitCell[k][1] = unitCell2[whk][1]
		unitCell[k][2] = unitCell2[whk][2]
	}

	// OrientationsEnergies[[k]][[1]] is the orientations available for
	// colour k, OrientationsEnergies[[k]][[2]] is the corresponding
	// adsorption energies
	orientationsEnergies := make([][][]float64, len(colors))
	for k := 0; k < len(colors); k++ {
		add := make([][]float64, 2)
		for i := 0; i < 2; i++ {
			add[i] = make([]float64, 0, len(unitCell2))
		}
		for i := 0; i < len(unitCell2); i++ {
			if int(unitCell2[i][0]) == colors[k] {
				add[0] = append(add[0], unitCell2[i][3])
				add[1] = append(add[1], unitCell2[i][4])
			}
		}
		orientationsEnergies[k] = add
	}

	// [R]
	// write.table(format(Lattice, digits=22, trim=T),
	//             file="Lattice.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	// write.table(t(as.matrix(Character-1)),
	//             file="Character.csv", sep=",", row.names=FALSE, col.names=FALSE, quote=F)
	//Lattice, character := LatticeGen(unitCell, LatticeVectors)
	Lattice := LoadFromCsvFile2Dim(path.Join(DATA_DIR, "Lattice.csv"), ',')
	character := LoadFromCsvFileInt(path.Join(DATA_DIR, "Character.csv"))

	// Make identical unit cell points?
	// 1 <-> 11, 4 <-> 8
	for i := 0; i < len(character); i++ {
		if character[i] == 0 {
			character[i] = 10
		} else if character[i] == 3 {
			character[i] = 7
		}
	}

	// Identify the unit cells by those which have character == central.point
	whC := Which(character, CentralPoint)

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
	// AdjSEQ := make([][][]int, Npower)
	// AdjSEQ[0] = Adj
	// for k := 1; k < Npower; k++ {
	// 	// Generate the power matrices
	// 	AdjSEQ[k] = MatrixMultiply(AdjSEQ[k-1], Adj)
	// }
	//
	// adjCuml := make([][][]int, Npower)
	// Copy2DimArray(&adjCuml[0], AdjSEQ[0])
	// for k := 1; k < Npower; k++ {
	// 	Copy2DimArray(&adjCuml[k], AdjSEQ[0])
	// 	for j := 1; j <= k; j++ {
	// 		MatrixAdd(adjCuml[k], AdjSEQ[j])
	// 	}
	// }

	// Generate all combinations of characters and corresponding
	// orientations. First space is empty - it gets filled in
	// Extension.Block
	//
	// charactersOrientations[][1] is index(0-base or 1-base)
	//
	chUnique := Unique(character)
	sort.Ints(chUnique)

	totalCombs := 0
	for _, v := range chUnique {
		totalCombs += len(orientationsEnergies[v][0])
	}
	charactersOrientations := Create2DimArrayInt(totalCombs, 3)
	cnt := 0
	for _, ch := range chUnique {
		for _, opos := range orientationsEnergies[ch][0] {
			charactersOrientations[cnt][1] = ch
			charactersOrientations[cnt][2] = int(math.Floor(opos + .5))
			cnt++
		}
	}

	Inp = &InitData{
		UnitCell:             unitCell,
		UnitCell2:            unitCell2,
		OrientationsEnergies: orientationsEnergies,
		UnitCellCoords:       unitCellCoords,
		//AdjCuml:                LoadFromCsvFileList(path.Join(DATA_DIR, "AdjCuml.csv")),
		AdjCuml:                [][][]int{{{0}}},
		Character:              character,
		ChUnique:               chUnique,
		CharactersOrientations: charactersOrientations,
		MoleculeCoordinates:    LoadMoleculeCoordinates(CCoords, HCoords, BrCoords),
	}

	SetZcoulomb()

	// Generate the list VARS of indices for the interacting Coulomb matrix
	if PcaRep {
		for k := 0; k < len(Inp.MoleculeCoordinates.All); k++ {
			for j := 0; j < len(Inp.MoleculeCoordinates.All); j++ {
				Vars = append(Vars, []int{k, j})
			}
		}
	}

	// Load KRLS objects
	LoadDataFromJSONFile(&KernelRegsRepLog, path.Join(DATA_DIR, "kernelregS_Rep_log.json"))
	LoadDataFromJSONFile(&KernelRegsAtt, path.Join(DATA_DIR, "kernelregS_Att.json"))

	// Load SVM objects
	LoadDataFromJSONFile(&SvmModel, path.Join(DATA_DIR, "svm_model.json"))
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
