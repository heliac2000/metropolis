//
// extension_reduction_probability_block.go
//

package functions

import . "../util"

// Compute the probability of the transiton Canon --> Canon_out
//
func ExtensionReductionProbabilityBlock(pcanon, ccanon [][]int, ocanon [][]float64,
	pcanon_out, ccanon_out [][]int, ocanon_out [][]float64) float64 {
	// Delete islands from these if they are isomorphic
	// Excluding zero cases
	l, lo := len(pcanon), len(pcanon_out)
	isomer := make([]int, lo-1)
	FillSlice(isomer, -1)
	for j := 0; j < lo-1; j++ {
		for k := 0; k < l-1; k++ {
			if Member(k, isomer) || len(pcanon[k]) != len(pcanon_out[j]) {
				continue
			}

			ijm, ijc, ijo := pcanon[k], ccanon[k], ocanon[k]
			ikm, ikc, iko := pcanon_out[j], ccanon_out[j], ocanon_out[j]
			if (len(ijm) == 1 && len(ikm) == 1 && ijm[0] == 0 && ikm[0] == 0) ||
				IsomorphIslandsBlock(ijm, ijc, ijo, ikm, ikc, iko) {
				isomer[j] = k
				break
			}
		}
	}

	coutInd := make([]int, 0, lo)
	for i, v := range isomer {
		if v == -1 {
			coutInd = append(coutInd, i)
		}
	}

	// Step 1. Do case of two new islands appearing
	switch len(coutInd) {
	case 2:
		return probBlockQabcd(
			pcanon, ccanon, ocanon, pcanon_out, ccanon_out, ocanon_out, coutInd[0], coutInd[1])
	case 1:
		return probBlockQabbd(
			pcanon, ccanon, ocanon, pcanon_out, ccanon_out, ocanon_out, coutInd[0])
	case 0:
		return probBlockQabba(
			pcanon, ccanon, ocanon, pcanon_out, ccanon_out, ocanon_out)
	}

	return 0.0
}

// Step 1. Do case of two new islands appearing
//
func probBlockQabcd(pcanon, ccanon [][]int, ocanon [][]float64,
	pcanon_out, ccanon_out [][]int, ocanon_out [][]float64, a, b int) float64 {
	// Permutation(l-1, 2) : 0-base
	l := len(pcanon)
	numGor := l * (l - 1)
	ch := make(chan float64, numGor)
	for i1 := 0; i1 < l; i1++ {
		for i2 := 0; i2 < l; i2++ {
			if i1 == i2 {
				continue
			}
			go func(i1, i2 int) {
				var zcoord [][]float64
				if len(pcanon[i2]) == 1 && pcanon[i2][0] == 0 {
					zcoord = [][]float64{{0, 0, 0}}
				} else {
					zcoord = CoordsIsland(pcanon[i2], ccanon[i2], ocanon[i2])
				}
				extb, lb := ExtensionBlock(pcanon[i2], zcoord)

				preda, creda, oreda := ReductionBlock(pcanon[i1], ccanon[i1], ocanon[i1])
				ret := QabcdBlock(pcanon[i1], pcanon[i2], ccanon[i2], ocanon[i2],
					pcanon_out[a], ccanon_out[a], ocanon_out[a], pcanon_out[b], ccanon_out[b], ocanon_out[b],
					preda, creda, oreda, pcanon, extb, lb, i1)
				ch <- ret
			}(i1, i2)
		}
	}

	qtestOut := 0.0
	for i := 0; i < numGor; i++ {
		qtestOut += <-ch
	}
	close(ch)

	return qtestOut
}

// Step 1. Do case of two new islands appearing
//
func probBlockQabbd(pcanon, ccanon [][]int, ocanon [][]float64,
	pcanon_out, ccanon_out [][]int, ocanon_out [][]float64, a int) float64 {
	// Permutation(l-1, 2) : 0-base
	l, qtestOut := len(pcanon), 0.0
	for i1 := 0; i1 < l; i1++ {
		for i2 := 0; i2 < l; i2++ {
			if i1 == i2 {
				continue
			}

			var zcoord [][]float64
			if len(pcanon[i2]) == 1 && pcanon[i2][0] == 0 {
				zcoord = [][]float64{{0, 0, 0}}
			} else {
				zcoord = CoordsIsland(pcanon[i2], ccanon[i2], ocanon[i2])
			}
			extb, lb := ExtensionBlock(pcanon[i2], zcoord)

			preda, creda, oreda := ReductionBlock(pcanon[i1], ccanon[i1], ocanon[i1])
			qtestOut += QabbdBlock(pcanon[i1], ccanon[i1], ocanon[i1],
				pcanon[i2], ccanon[i2], ocanon[i2], pcanon_out[a], ccanon_out[a], ocanon_out[a],
				preda, creda, oreda, pcanon, extb, lb, i1) +
				Qab0dBlock(pcanon[i1], ccanon[i1], ocanon[i1],
					pcanon[i2], ccanon[i2], ocanon[i2], pcanon_out[a], ccanon_out[a], ocanon_out[a],
					preda, creda, oreda, pcanon, extb, lb, i1)
		}
	}

	return qtestOut
}

// Step 1. Do case of two new islands appearing
//
func probBlockQabba(pcanon, ccanon [][]int, ocanon [][]float64,
	pcanon_out, ccanon_out [][]int, ocanon_out [][]float64) float64 {
	// Permutation(l-1, 2) : 0-base
	l, qtestOut := len(pcanon), 0.0
	for i1 := 0; i1 < l; i1++ {
		for i2 := 0; i2 < l; i2++ {
			if i1 == i2 {
				continue
			}

			var zcoord [][]float64
			if len(pcanon[i2]) == 1 && pcanon[i2][0] == 0 {
				zcoord = [][]float64{{0, 0, 0}}
			} else {
				zcoord = CoordsIsland(pcanon[i2], ccanon[i2], ocanon[i2])
			}
			extb, lb := ExtensionBlock(pcanon[i2], zcoord)

			preda, creda, oreda := ReductionBlock(pcanon[i1], ccanon[i1], ocanon[i1])
			qtestOut += QabbaBlock(pcanon[i1], ccanon[i1], ocanon[i1],
				pcanon[i2], ccanon[i2], ocanon[i2], preda, creda, oreda, pcanon, extb, lb, i1) +
				Qab0aBlock(pcanon[i1], ccanon[i1], ocanon[i1],
					pcanon[i2], ccanon[i2], ocanon[i2], preda, creda, oreda, pcanon, extb, lb, i1)
		}
	}

	return qtestOut
}
