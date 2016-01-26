//
// energy_canonical.go
//

package functions

// Get the total energy of a canonical configuration (CtestM[[1]] =
// positions, CtestM[[2]] = orientations)
//
func EnergyCanonical(pos, chr [][]int, ori [][]float64) float64 {
	ene := 0.0

	// Don't include the final zero term
	for k := 0; k < len(pos)-1; k++ {
		// Edit as appropriate!
		ene += EnergyIslandPCA(pos[k], chr[k], ori[k])
	}

	return ene
}
