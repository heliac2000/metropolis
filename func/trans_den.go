//
// trans_den.go
//

package functions

import . "../util"

// Determine the translational degeneracy factors of Ctemp (factors A1
// to Ar in eq (63), pg. 284 of notes). Make sure that Ctemp has no
// zero entries!
func TransDen(pos, chr [][]int, ori [][]float64) []int {
	// Classify objects in Ctemp into equivalence classes according to
	// their translational symmetry
	l := len(pos)
	eclass := make([]int, l)
	// NOTICE: R is 1-base index, golang is 0-base.
	for i := 1; i < l; i++ { // eclass[0] == 0
		eclass[i] = -1
	}

	classified, unclassified := make([]int, 0, l), make([]int, 0, l)
	// Indices of islands that have been classified
	classified = append(classified, 0)
	// Indices of islands that have not been classified
	unclassified = append(unclassified, SeqInt(1, len(eclass)-1, 1)...)

	for _, k := range unclassified {
		eclass[k] = k
		// Check if this island has already appeared
		for _, j := range classified {
			cjClass := eclass[j]
			if len(pos[k]) == len(pos[j]) &&
				TranslationIdentical(pos[k], chr[k], ori[k], pos[j], chr[j], ori[j]) {
				eclass[k] = cjClass
				break
			}
		}

		classified = nil
		for i, v := range eclass {
			if v != -1 {
				// Indices of islands that have been classified
				classified = append(classified, i)
			}
		}
	}

	uniqueEclass := Unique(eclass)
	factors := make([]int, len(uniqueEclass))
	for i, k := range uniqueEclass {
		factors[i] = CountItems(k, eclass)
	}

	return factors
}
