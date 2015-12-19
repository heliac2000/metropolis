//
// lattice_gen.go
//

package functions

import . "../util"

// Generate a lattice based on the unit cell
//
func LatticeGen(unitCell, latticeVectors [][]float64) ([][]float64, []int) {
	// Number of lattice points in the unit cell
	nlp := len(unitCell)

	// Generate the coordinates of the lattice from the unit cell
	latticeCoords := Create2DimArrayFloat(nlp, 2, (nlp+1)*Nrepeat*Nrepeat/4)

	// horizontal/vertical coordinate
	for k := 0; k < nlp; k++ {
		latticeCoords[k][0], latticeCoords[k][1] = unitCell[k][0], unitCell[k][1]
	}

	// Specify which unit cell point the lattice point corresponds to
	character := make([]int, 0, (nlp+1)*Nrepeat*Nrepeat/4)
	seq := func(n int) []int {
		seq := make([]int, n)
		for i := 0; i < n; i++ {
			//seq[i] = i + 1
			seq[i] = i
		}
		return seq
	}(nlp)
	character = append(character, seq...)

	var unitCellC [][]float64
	avec, bvec := latticeVectors[0], latticeVectors[1]
	Copy2DimArray(&unitCellC, latticeCoords)
	for k := 0; k < Nrepeat/2; k++ {
		var latticeTemp [][]float64
		for j := 0; j < Nrepeat/2; j++ {
			Copy2DimArray(&latticeTemp, unitCellC)
			for h := 0; h < nlp; h++ {
				latticeTemp[h][0] += float64(k)*avec[0] + float64(j)*bvec[0]
				latticeTemp[h][1] += float64(k)*avec[1] + float64(j)*bvec[1]
			}
			latticeCoords = append(latticeCoords, latticeTemp...)
			character = append(character, seq...)
		}
	}

	// Matrix of lattice coordinates
	lattice, remv := MatrixTidy(latticeCoords)
	char := make([]int, 0, len(character))
	for i := 0; i < len(character); i++ {
		if remv[i] == 0 {
			char = append(char, character[i])
		}
	}

	return lattice, char
}
