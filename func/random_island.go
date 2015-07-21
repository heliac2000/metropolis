//
// random_island.go
//

package functions

// Generate a random island with n occupied unit cells
//
func RandomIslandUnitCell(n int, adjCj [][]int,
	unitCell [][]float64, chUnique []int) ([]int, []int, []float64) {
	islandP := make([]int, 0, n)
	islandP = append(islandP, UCcenter)
	if n > 1 {
		for k := 1; k < n; k++ {
			surrIsland := SurrAdj(islandP, adjCj)
			islandP = append(islandP, surrIsland[Rnd.Intn(len(surrIsland))])
		}
	}

	// Generate the orientations

	// Output characters of the adsorbed molecules
	// Output orientations of the adsorbed molecules
	islandC, islandO := make([]int, 0, n), make([]float64, 0, n)
	for k := 0; k < n; k++ {
		islandC = append(islandC, chUnique[Rnd.Intn(len(chUnique))])
		islandO = append(islandO, unitCell[islandC[k]][4+Rnd.Intn(3)])
	}

	return islandP, islandC, islandO
}
