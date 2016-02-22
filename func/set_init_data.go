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
func SetInitData(dataDir string) {
	unitCell2 := LoadFromCsvFile2Dim(path.Join(dataDir, UnitCell2File), ',')
	//latticeVectors := LoadFromCsvFile2Dim(path.Join(dataDir, LatticeVectorsFile), ',')

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
	//Lattice, character := LatticeGen(unitCell, latticeVectors)
	lattice := LoadFromCsvFile2Dim(path.Join(dataDir, LatticeFile), ',')
	character := LoadFromCsvFileInt(path.Join(dataDir, CharacterFile))

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
		copy(unitCellCoords[k], lattice[whC[k]])
	}

	// Make a sequence of power matrices
	//
	// Moves, Adj := Create2DimArrayFloat(4, 2), Create2DimArrayInt(nUC, nUC)
	// avec, bvec := latticeVectors[0], latticeVectors[1]
	// for j := 0; j < nUC; j++ {
	// 	for i := 0; i < 4; i++ {
	// 		copy(Moves[i], unitCellCoords[j])
	// 	}
	// 	Moves[0][0] += avec[0]
	// 	Moves[1][0] -= avec[0]
	// 	Moves[2][0] += bvec[0]
	// 	Moves[2][1] += bvec[1]
	// 	Moves[3][0] -= bvec[0]
	// 	Moves[3][1] -= bvec[1]
	//
	// 	_, surrj := GetKnnx(unitCellCoords, Moves, 1)
	// 	surr := make([]int, len(surrj))
	// 	for i := 0; i < len(surrj); i++ {
	// 		surr[i] = surrj[i][0]
	// 	}
	// 	surr = Unique(surr)
	//
	// 	for i := 0; i < len(surr); i++ {
	// 		Adj[j][surr[i]] = 1
	// 	}
	// 	Adj[j][j] = 0
	// }
	//
	// adjSEQ := make([][][]int, Npower)
	// adjSEQ[0] = Adj
	// for k := 1; k < Npower; k++ {
	// 	// Generate the power matrices
	// 	adjSEQ[k] = MatrixMultiply(adjSEQ[k-1], Adj)
	// }
	//
	// adjCuml := make([][][]int, Npower)
	// adjCuml[0] = Copy2DimArrayInt(adjSEQ[0])
	// var wg sync.WaitGroup
	// wg.Add(Npower - 1)
	// for k := 1; k < Npower; k++ {
	// 	go func(k int) {
	// 		adjCuml[k] = Copy2DimArrayInt(adjSEQ[0])
	// 		for j := 1; j <= k; j++ {
	// 			MatrixAdd(adjCuml[k], adjSEQ[j])
	// 		}
	// 		wg.Done()
	// 	}(k)
	// }
	// wg.Wait()

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
		UnitCell:               unitCell,
		UnitCell2:              unitCell2,
		OrientationsEnergies:   orientationsEnergies,
		UnitCellCoords:         unitCellCoords,
		AdjCuml:                LoadFromCsvFileList(path.Join(dataDir, AdjCumlFile)),
		Character:              character,
		ChUnique:               chUnique,
		CharactersOrientations: charactersOrientations,
		MoleculeCoordinates:    LoadMoleculeCoordinates(dataDir, CCoords, HCoords, BrCoords),
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
	LoadDataFromJSONFile(&KernelRegsRepLog, path.Join(dataDir, KernelRegsRepLogFile))
	KernelRegsRepLog.ColMeansX = ColMeans(KernelRegsRepLog.X)
	KernelRegsRepLog.ColMeansY = ColMeans(KernelRegsRepLog.Y)
	KernelRegsRepLog.ColSdX = ColSd(KernelRegsRepLog.X)
	KernelRegsRepLog.ColSdY = ColSd(KernelRegsRepLog.Y)
	KernelRegsRepLog.ScaleX = Scale(KernelRegsRepLog.X, KernelRegsRepLog.ColMeansX, KernelRegsRepLog.ColSdX)

	LoadDataFromJSONFile(&KernelRegsAtt, path.Join(dataDir, KernelRegsAttFile))
	KernelRegsAtt.ColMeansX = ColMeans(KernelRegsAtt.X)
	KernelRegsAtt.ColMeansY = ColMeans(KernelRegsAtt.Y)
	KernelRegsAtt.ColSdX = ColSd(KernelRegsAtt.X)
	KernelRegsAtt.ColSdY = ColSd(KernelRegsAtt.Y)
	KernelRegsAtt.ScaleX = Scale(KernelRegsAtt.X, KernelRegsAtt.ColMeansX, KernelRegsAtt.ColSdX)

	// Load SVM objects
	LoadDataFromJSONFile(&SvmModel, path.Join(dataDir, SvmModelFile))
	LoadDataFromJSONFile(&SvmModelOp, path.Join(dataDir, SvmModelOpFile))
	LoadDataFromJSONFile(&SvmModelNzp, path.Join(dataDir, SvmModelNzpFile))
	LoadDataFromJSONFile(&SvmModelUsp, path.Join(dataDir, SvmModelUspFile))

	// Load PrComp objects
	LoadDataFromJSONFile(&XeigPc, path.Join(dataDir, XeigPcFile))
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
