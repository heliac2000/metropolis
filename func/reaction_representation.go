//
// reaction_representation.go
//

package functions

import (
	"sort"

	. "../util"
)

// Try out reaction representation of Ci --> Cj transformation
// Reactant is Ci, Product is Cj
//
func ReactionRepresentation(prct, crct [][]int, orct [][]float64,
	ppdt, cpdt [][]int, opdt [][]float64) ([][]int, [][]int, [][]float64, []int, []int, []int) {

	cPrct, cPpdt := CopyArrayInt(prct), CopyArrayInt(ppdt)
	cCrct, cCpdt := CopyArrayInt(crct), CopyArrayInt(cpdt)
	cOrct, cOpdt := CopyArrayFloat(orct), CopyArrayFloat(opdt)

	ptot := make([][]int, len(prct)+len(ppdt))
	ctot := make([][]int, len(crct)+len(cpdt))
	otot := make([][]float64, len(orct)+len(opdt))

	for i := 0; i < len(cPrct); i++ {
		ptot[i] = make([]int, len(cPrct[i]))
		copy(ptot[i], cPrct[i])
	}
	for i, l := 0, len(cPrct); i < len(cPpdt); i++ {
		ptot[i+l] = make([]int, len(cPpdt[i]))
		copy(ptot[i+l], cPpdt[i])
	}

	for i := 0; i < len(cCrct); i++ {
		ctot[i] = make([]int, len(cCrct[i]))
		copy(ctot[i], cCrct[i])
	}
	for i, l := 0, len(cCrct); i < len(cCpdt); i++ {
		ctot[i+l] = make([]int, len(cCpdt[i]))
		copy(ctot[i+l], cCpdt[i])
	}

	for i := 0; i < len(cOrct); i++ {
		otot[i] = make([]float64, len(cOrct[i]))
		copy(otot[i], cOrct[i])
	}
	for i, l := 0, len(cOrct); i < len(cOpdt); i++ {
		otot[i+l] = make([]float64, len(cOpdt[i]))
		copy(otot[i+l], cOpdt[i])
	}

	// Rapid check for isomorphism between islands

	// Number of islands to consider
	nI := len(ptot)
	// Assign a label to each island according to its equivalence class
	labels, done := make([]int, nI), make([]int, nI)

	for k := 0; k < nI-1; k++ {
		if done[k] == 0 {
			done[k], labels[k] = 1, k
			for j := k + 1; j < nI; j++ {
				// [R code] if(length(Xk) != length(Xj)) dont = TRUE
				// lenght(Xk) and length(Xj) are always 1.
				if (ptot[k][0] == 0 && ptot[j][0] == 0) ||
					IsomorphIslandsBlock(ptot[k], ctot[k], otot[k], ptot[j], ctot[j], otot[j]) {
					labels[j], done[j] = k, 1
				}
			}
		}
	}

	// Identifies the unique classes
	labUnq := Unique(labels)
	sort.Ints(labUnq)

	l := len(labUnq)
	pil := Create2DimArrayInt(l, len(ptot[0]))
	cil := Create2DimArrayInt(l, len(ctot[0]))
	oil := Create2DimArrayFloat(l, len(otot[0]))
	for i, k := range labUnq {
		pil[i], cil[i], oil[i] = ptot[k], ctot[k], otot[k]
	}

	// lab1:Labels for reactant side, lab2:Labels for product side
	lab1, lab2 := CopyVectorInt(labels[:len(prct)]), CopyVectorInt(labels[len(prct):])
	// Coefficients for reactant side, Coefficients for product side
	coeffsRct, coeffsPdt := make([]int, 0, l), make([]int, 0, l)
	for _, k := range labUnq {
		coeffsRct = append(coeffsRct, CountItems(k, lab1))
		coeffsPdt = append(coeffsPdt, CountItems(k, lab2))
	}

	diff := make([]int, l)
	for i := 0; i < l; i++ {
		diff[i] = coeffsPdt[i] - coeffsRct[i]
	}

	return pil, cil, oil, coeffsRct, coeffsPdt, diff
}
