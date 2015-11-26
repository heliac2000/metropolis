//
// energy_island.go
//

package functions

// Calculate energy of an island (Island[[1]] = positions, Island[[2]]
// = characters, Island[[2]] = orientations)
func EnergyIsland(pos, chr []int, ori []float64) float64 {
	ene := 0.0

	// Adsorption energy
	for k := 0; k < len(pos); k++ {
		ch := Inp.Character[chr[k]]
		o1 := ori[k]
		ao := Inp.UnitCell2[ch]

		for i := 5; i < 8; i++ {
			if ao[i] == o1 {
				ene += ao[i+3]
				break
			}
		}
	}

	eneInt := 0.0
	if len(pos) > 1 {
		for k := 0; k < len(pos)-1; k++ {
			for j := k + 1; j < len(pos); j++ {
				_, ep := EnergyPair(pos[k], pos[j], chr[k], chr[j], ori[k], ori[j])
				eneInt += ep
			}
		}
	}

	return ene + eneInt
}
