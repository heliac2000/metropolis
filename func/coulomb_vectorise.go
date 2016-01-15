//
// coulomb_vectorise.go
//

package functions

// Vectorise the interacting Coulomb matrix CMtest
//
func CoulombVectorise(cmTest [][]float64) []float64 {
	cmVec := make([]float64, len(Vars))

	for i, v := range Vars {
		cmVec[i] = cmTest[v[0]][v[1]]
	}

	return cmVec
}
