//
// canonical_gen.go
//

package functions

// Functions for canonical representations
//
// Generate a random canonical representation of random islands. For
// now, make all sizes = 1
//
func CanonicalGen() ([][]int, [][]int, [][]float64) {
	// Need to make this better - a truly uniform random initial
	// condition please!

	xout, cout, oout :=
		make([][]int, Nmolec), make([][]int, Nmolec), make([][]float64, Nmolec)
	for k := 0; k < Nmolec; k++ {
		xout[k], cout[k], oout[k] = RandomIslandUnitCell(1)
	}

	return xout, cout, oout
}
