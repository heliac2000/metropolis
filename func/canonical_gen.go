//
// canonical_gen.go
//

package functions

// Functions for canonical representations
//
// Generate a random canonical representation of random islands. For
// now, make all sizes = 1
//
func CanonicalGen() ([][][]int, [][][]int, [][][]float64) {
	xout, cout, oout :=
		make([][][]int, Nmolec), make([][][]int, Nmolec), make([][][]float64, Nmolec)

	for k := 0; k < Nmolec; k++ {
		isn1, isn2, isn3 := RandomIslandUnitCell(1)
		xout[k], cout[k], oout[k] = [][]int{isn1}, [][]int{isn2}, [][]float64{isn3}
	}

	return xout, cout, oout
}
