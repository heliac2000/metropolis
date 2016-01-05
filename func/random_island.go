//
// random_island.go
//

package functions

// Generate a random island with n occupied unit cells
//
func RandomIslandUnitCell(n int) ([]int, []int, []float64) {
	islandP := make([]int, 0, n)
	islandP = append(islandP, UCcenter)
	if n > 1 {
		for k := 1; k < n; k++ {
			surrIsland := SurrAdj(islandP, Inp.AdjCuml[Npower-1])
			islandP = append(islandP, surrIsland[Rnd.Intn(len(surrIsland))])
		}
	}

	// Generate the orientations

	// Output characters of the adsorbed molecules
	// Output orientations of the adsorbed molecules
	islandC, islandO := make([]int, 0, n), make([]float64, 0, n)
	for k := 0; k < n; k++ {
		islandC = append(islandC, Inp.ChUnique[Rnd.Intn(len(Inp.ChUnique))])
		oAvail := Inp.OrientationsEnergies[islandC[k]][0]
		islandO = append(islandO, oAvail[Rnd.Intn(len(oAvail))])
	}

	return islandP, islandC, islandO
}
