//
// lattice_gen.go
//

package functions

import . "../util"

// Generate a lattice based on the unit cell
//
func LatticeGen(UnitCell, LatticeVectors [][]float64) ([][]float64, []int) {
	// Number of lattice points in the unit cell
	nlp := len(UnitCell)

	// Generate the coordinates of the lattice from the unit cell
	LatticeCoords := Create2DimArray(nlp, 2, (nlp+1)*Nrepeat*Nrepeat/4)

	// horizontal/vertical coordinate
	for k := 0; k < nlp; k++ {
		LatticeCoords[k][0], LatticeCoords[k][1] = UnitCell[k][0], UnitCell[k][1]
	}

	// Specify which unit cell point the lattice point corresponds to
	Character := make([]int, nlp, (nlp+1)*Nrepeat*Nrepeat/4)
	for i := 0; i < nlp; i++ {
		Character[i] = i + 1
	}

	avec, bvec := LatticeVectors[0], LatticeVectors[1]
	UnitCellC := Copy2DimArray(LatticeCoords)
	for k := 0; k < Nrepeat/2; k++ {
		for j := 0; j < Nrepeat/2; j++ {
			LatticeTemp := Copy2DimArray(UnitCellC)
			for h := 0; h < nlp; h++ {
				LatticeTemp[h][0] += float64(k)*avec[0] + float64(j)*bvec[0]
				LatticeTemp[h][1] += float64(k)*avec[1] + float64(j)*bvec[1]
			}
			LatticeCoords = append(LatticeCoords, LatticeTemp...)
			for i := 0; i < nlp; i++ {
				Character = append(Character, i+1)
			}
		}
	}

	// Matrix of lattice coordinates
	Out, Remv := MatrixTidy(LatticeCoords)

	return Out, Remv
}
