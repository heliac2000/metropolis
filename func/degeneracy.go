//
// degeneracy.go
//

package functions

import . "../util"

// Compute the degeneracy factor (from equation 63, page 284 of notes)
//
func Degeneracy(pos, chr [][]int, ori [][]float64) float64 {
	// Determine the number of non-zero islands in Ctest. Important for
	// Ctest to be ordered.
	clength := 0
	for k := 0; k < len(pos); k++ {
		sum := 0
		for _, v := range pos[k] {
			sum += v
		}
		if sum == 0 {
			clength = k
			break
		}
	}

	// Remove zeroes from canonical representation.
	// rot: Store all rotated islands, rotInd: Unique indices
	rotInd, rot := make([][]int, clength), make([][]int, clength)
	for k := 0; k < clength; k++ {
		rotInd[k], rot[k] = UniqueOrientations(pos[k])
	}

	// Now, determine all unique combinations of indices from Rot
	rotCombs, degF := ExpandGrid(rotInd), 0.0
	for _, indsk := range rotCombs {
		ppos := make([][]int, clength)
		for j := 0; j < clength; j++ {
			if ind := rotInd[j][indsk[j]]; ind == 0 {
				ppos[j] = pos[j]
			} else {
				ppos[j] = rot[j]
			}
		}
		factors := TransDen(ppos, chr[:clength], ori[:clength])
		addF := 1.0
		for _, v := range factors {
			addF *= Factorial(v)
		}
		degF += 1.0 / addF
	}

	// Degeneracy factor
	return Factorial(clength) * ApproXnCr(Nuc*Nuc, clength) * degF
}
