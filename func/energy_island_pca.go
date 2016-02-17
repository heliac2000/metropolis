//
// energy_island_pca.go
//

package functions

import "math"

// Calculate energy of an island (Island[[1]] = positions, Island[[2]]
// = characters, Island[[2]] = orientations). Only use this if PCA_REP
// = TRUE
//
func EnergyIslandPCA(pos, chr []int, ori []float64) float64 {
	// Adsorption energy
	ene := 0.0
	for k := 0; k < len(pos); k++ {
		// Incase new UnitCell2 file format is used
		for _, u := range Inp.UnitCell2 {
			// u[0]: color, u[3]: angle, u[4]: energy
			if int(u[0]) == Inp.Character[chr[k]]+1 && math.Abs(u[3]-ori[k]) < Detect {
				ene += u[4]
			}
		}
	}

	if len(pos) <= 1 {
		return ene
	}

	// Interaction energy
	for k := 0; k < len(pos)-1; k++ {
		for j := k + 1; j < len(pos); j++ {
			ene += EnergyPairReduce(pos[k], pos[j], chr[k], chr[j], ori[k], ori[j])
		}
	}

	return ene
}
