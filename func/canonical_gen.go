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
	sizes := make([]int, Nmolec)
	for k, sum := 0, 0; k < Nmolec; k++ {
		if k == Nmolec-1 {
			sizes[k] = Nmolec - sum
		} else {
			sizes[k] = Rnd.Intn(Nmolec - sum + 1)
			sum += sizes[k]
			if sum == Nmolec {
				break
			}
		}
	}

	xout, cout, oout :=
		make([][]int, Nmolec), make([][]int, Nmolec), make([][]float64, Nmolec)
	for k := 0; k < Nmolec; k++ {
		if sizes[k] > 0 {
			xout[k], cout[k], oout[k] = RandomIslandUnitCell(sizes[k])
		} else {
			xout[k], cout[k], oout[k] = []int{0}, []int{0}, []float64{0}
		}
	}

	return xout, cout, oout
}
