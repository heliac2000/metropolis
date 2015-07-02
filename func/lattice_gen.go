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
	LatticeCoords := Create2DimArray(float64(0), nlp, 2, (nlp+1)*Nrepeat*Nrepeat/4).([][]float64)

	// horizontal/vertical coordinate
	for k := 0; k < nlp; k++ {
		LatticeCoords[k][0], LatticeCoords[k][1] = UnitCell[k][0], UnitCell[k][1]
	}

	// Specify which unit cell point the lattice point corresponds to
	Character := make([]int, 0, (nlp+1)*Nrepeat*Nrepeat/4)
	seq := func(n int) []int {
		seq := make([]int, n)
		for i := 0; i < n; i++ {
			seq[i] = i + 1
		}
		return seq
	}(nlp)
	Character = append(Character, seq...)

	avec, bvec := LatticeVectors[0], LatticeVectors[1]
	UnitCellC := Copy2DimArray(LatticeCoords)
	for k := 0; k < Nrepeat/2; k++ {
		for j := 0; j < Nrepeat/2; j++ {
			LatticeTemp := Copy2DimArray(UnitCellC).([][]float64)
			for h := 0; h < nlp; h++ {
				LatticeTemp[h][0] += float64(k)*avec[0] + float64(j)*bvec[0]
				LatticeTemp[h][1] += float64(k)*avec[1] + float64(j)*bvec[1]
			}
			LatticeCoords = append(LatticeCoords, LatticeTemp...)
			Character = append(Character, seq...)
		}
	}

	// Matrix of lattice coordinates
	Lattice, Remv := MatrixTidy(LatticeCoords)
	char := make([]int, 0, len(Character))
	for i := 0; i < len(Character); i++ {
		if Remv[i] == 0 {
			char = append(char, Character[i])
		}
	}

	return Lattice, char
}
