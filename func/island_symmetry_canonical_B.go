//
// island_symmetry_canonical_B.go
//

package functions

// Compute the product of the symmetry factors for molecules
//
func IslandSymmetryCanonicalB(pos [][]int) float64 {
	fact := 1.0

	for k := 0; k < len(pos); k++ {
		// Factor of two arises because this is a two-fold lattice.
		fact *= 2 / IslandSymmetryBlock(pos[k])
	}

	return fact
}
