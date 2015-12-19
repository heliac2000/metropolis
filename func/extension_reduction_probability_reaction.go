//
// extension_reduction_probability_reaction.go
//

package functions

import (
	"reflect"
	"strconv"

	. "../util"
)

type eRPRFunction struct {
	prct, crct                [][]int
	orct                      [][]float64
	ppdt, cpdt                [][]int
	opdt                      [][]float64
	pil, cil                  [][]int
	oil                       [][]float64
	rCoeffs, pCoeffs, dCoeffs []int
	rclass                    int
}

// Probability terms
// Take care when calling the probability factors q.
// qadcd.Block returns qab,cd + qab,dc
// qabbd.Block returns qab,bd + qab,da
// qab0d.Block returns qab,0d only.
// qabba.Block returns qab,ba only
// qab0a.Block returns qab,0a only

// Compute the probability of the transiton Canon --> Canon_out based
// on reaction representation.
//
func ExtensionReductionProbabilityReaction(prct, crct [][]int, orct [][]float64,
	ppdt, cpdt [][]int, opdt [][]float64) float64 {
	pil, cil, oil, rCoeffs, pCoeffs, dCoeffs :=
		ReactionRepresentation(prct, crct, orct, ppdt, cpdt, opdt)

	rclass := ReactionClassID(dCoeffs)
	if rclass == REACT_UNKNOWN {
		return 0.0
	}

	return (reflect.ValueOf(&eRPRFunction{
		prct, crct, orct, ppdt, cpdt, opdt, pil, cil, oil, rCoeffs, pCoeffs, dCoeffs, rclass,
	}).MethodByName("ERPRClass" + strconv.Itoa(rclass+1)).Call(nil))[0].Float()
}

// REACT_CLASS1
//
func (e *eRPRFunction) ERPRClass1() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -1), Which(e.dCoeffs, 1)
	whR1, whR2, whP1, whP2 := whR[0], whR[1], whP[0], whP[1]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	pR2, cR2, oR2 := e.pil[whR2], e.cil[whR2], e.oil[whR2]
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	pP2, cP2, oP2 := e.pil[whP2], e.cil[whP2], e.oil[whP2]
	// Coefficients for the islands R1 and R2 on the reactant side
	nR1, nR2 := float64(e.rCoeffs[whR1]), float64(e.rCoeffs[whR2])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	pRedR2, cRedR2, oRedR2 := ReductionBlock(pR2, cR2, oR2)
	extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))
	extR2, lR2 := ExtensionBlock(pR2, CoordsIsland(pR2, cR2, oR2))

	return nR1*nR2*QabcdBlock(
		pR1, pR2, cR2, oR2, pP1, cP1, oP1, pP2, cP2, oP2,
		pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1) +
		nR1*nR2*QabcdBlock(
			pR2, pR1, cR1, oR1, pP1, cP1, oP1, pP2, cP2, oP2,
			pRedR2, cRedR2, oRedR2, e.prct, extR1, lR1, 1)
}

// REACT_CLASS2
//
func (e *eRPRFunction) ERPRClass2() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -1), Which(e.dCoeffs, 2)
	whR1, whR2, whP1 := whR[0], whR[1], whP[0]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	pR2, cR2, oR2 := e.pil[whR2], e.cil[whR2], e.oil[whR2]
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	// Coefficients for the islands R1 and R2 on the reactant side
	nR1, nR2 := float64(e.rCoeffs[whR1]), float64(e.rCoeffs[whR2])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	pRedR2, cRedR2, oRedR2 := ReductionBlock(pR2, cR2, oR2)
	extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))
	extR2, lR2 := ExtensionBlock(pR2, CoordsIsland(pR2, cR2, oR2))

	return nR1*nR2*QabcdBlock(
		pR1, pR2, cR2, oR2, pP1, cP1, oP1, pP1, cP1, oP1,
		pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1) +
		nR1*nR2*QabcdBlock(
			pR2, pR1, cR1, oR1, pP1, cP1, oP1, pP1, cP1, oP1,
			pRedR2, cRedR2, oRedR2, e.prct, extR1, lR1, 1)
}

// REACT_CLASS3
//
func (e *eRPRFunction) ERPRClass3() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -1), Which(e.dCoeffs, 1)
	whR1, whR2, whP1 := whR[0], whR[1], whP[0]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	pR2, cR2, oR2 := e.pil[whR2], e.cil[whR2], e.oil[whR2]
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	// Coefficients for the islands R1 and R2 on the reactant side
	nR1, nR2 := float64(e.rCoeffs[whR1]), float64(e.rCoeffs[whR2])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	pRedR2, cRedR2, oRedR2 := ReductionBlock(pR2, cR2, oR2)
	extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))
	extR2, lR2 := ExtensionBlock(pR2, CoordsIsland(pR2, cR2, oR2))

	return nR1*nR2*Qab0dBlock(
		pR1, cR1, oR1, pR2, cR2, oR2, pP1, cP1, oP1,
		pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1) +
		nR1*nR2*Qab0dBlock(
			pR2, cR2, oR2, pR1, cR1, oR1, pP1, cP1, oP1,
			pRedR2, cRedR2, oRedR2, e.prct, extR1, lR1, 1)
}

// REACT_CLASS4
//
func (e *eRPRFunction) ERPRClass4() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -2), Which(e.dCoeffs, 1)
	whR1, whP1, whP2 := whR[0], whP[0], whP[1]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	pP2, cP2, oP2 := e.pil[whP2], e.cil[whP2], e.oil[whP2]
	// Coefficients for the islands R1 and R2 on the reactant side
	nR1 := float64(e.rCoeffs[whR1])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))

	return nR1 * nR1 * QabcdBlock(
		pR1, pR1, cR1, oR1, pP1, cP1, oP1, pP2, cP2, oP2,
		pRedR1, cRedR1, oRedR1, e.prct, extR1, lR1, 1)
}

// REACT_CLASS5
//
func (e *eRPRFunction) ERPRClass5() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -2), Which(e.dCoeffs, 1)
	whR1, whP1 := whR[0], whP[0]
	// R1: Reactant island 1, P1: Product island 1
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	// Coefficients for the islands R1 on the reactant side
	nR1 := float64(e.rCoeffs[whR1])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))

	return nR1 * nR1 * Qab0dBlock(
		pR1, cR1, oR1, pR1, cR1, oR1, pP1, cP1, oP1,
		pRedR1, cRedR1, oRedR1, e.prct, extR1, lR1, 1)
}

// REACT_CLASS6
//
func (e *eRPRFunction) ERPRClass6() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -1), Which(e.dCoeffs, 1)
	whR1, whP1, whP2 := whR[0], whP[0], whP[1]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	pR2, cR2, oR2 := []int{0}, []int{0}, []float64{0}
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	pP2, cP2, oP2 := e.pil[whP2], e.cil[whP2], e.oil[whP2]
	// Coefficients for the islands R1 on the reactant side
	nR1 := float64(e.rCoeffs[whR1])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	extR2, lR2 := ExtensionBlock(pR2, [][]float64{{0, 0, 0}})

	return nR1 * QabcdBlock(
		pR1, pR2, cR2, oR2, pP1, cP1, oP1, pP2, cP2, oP2,
		pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1)
}

// REACT_CLASS7
//
func (e *eRPRFunction) ERPRClass7() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -1), Which(e.dCoeffs, 2)
	whR1, whP1 := whR[0], whP[0]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	pR2, cR2, oR2 := []int{0}, []int{0}, []float64{0}
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	// Coefficients for the islands R1 on the reactant side
	nR1 := float64(e.rCoeffs[whR1])

	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)

	// Generate the zero island
	extR2, lR2 := ExtensionBlock(pR2, [][]float64{{0, 0, 0}})

	return nR1 * QabcdBlock(
		pR1, pR2, cR2, oR2, pP1, cP1, oP1, pP1, cP1, oP1,
		pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1)
}

// REACT_CLASS8
//
func (e *eRPRFunction) ERPRClass8() float64 {
	// whR: Identify what was affected during the ER transformation
	// whP: Identify what was produced during the ER transformation
	whR, whP := Which(e.dCoeffs, -1), Which(e.dCoeffs, 1)
	whR1, whP1 := whR[0], whP[0]
	// R1: Reactant island 1, R2: Reactant island 2
	pR1, cR1, oR1 := e.pil[whR1], e.cil[whR1], e.oil[whR1]
	// P1: Product island 1, P2: Product island 2
	pP1, cP1, oP1 := e.pil[whP1], e.cil[whP1], e.oil[whP1]
	// Coefficients for the islands R1 on the reactant side
	nR1 := float64(e.rCoeffs[whR1])
	pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
	extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))

	qTestTot := 0.0
	for k := 0; k < len(e.pil); k++ {
		pR2, cR2, oR2 := e.pil[k], e.cil[k], e.oil[k]
		// Zero island condition
		if !(pR2[0] == 0 && cR2[0] == 0) &&
			!IsomorphIslandsBlock(pR1, cR1, oR1, pR2, cR2, oR2) {
			nR2 := float64(e.rCoeffs[k])
			extR2, lR2 := ExtensionBlock(pR2, CoordsIsland(pR2, cR2, oR2))
			pRedR2, cRedR2, oRedR2 := ReductionBlock(pR2, cR2, oR2)

			qTestTot += nR1*nR2*QabcdBlock(
				pR2, pR1, cR1, oR1, pP1, cP1, oR2, pR2, cR2, oR2,
				pRedR2, cRedR2, oRedR2, e.prct, extR1, lR1, 1) +
				nR1*nR2*QabcdBlock(
					pR1, pR2, cR2, oR2, pR2, cR2, oR2, pP1, cP1, oP1,
					pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1)
			qTestTot *= 2
		}
	}

	// Include the zero block correction
	pR2, cR2, oR2 := []int{0}, []int{0}, []float64{0}
	extR2, lR2 := ExtensionBlock(pR2, [][]float64{{0, 0, 0}})

	return qTestTot + QabcdBlock(
		pR1, pR2, cR2, oR2, pR2, cR2, oR2, pP1, cP1, oP1,
		pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1)
}

// REACT_CLASS9
//
func (e *eRPRFunction) ERPRClass9() float64 {
	// Create the zero block
	pZb, cZb, oZb := []int{0}, []int{0}, []float64{0}
	extZb, lZb := ExtensionBlock(pZb, [][]float64{{0, 0, 0}})
	l := len(e.pil)

	// Compute the diagonal term
	qTestTot := 0.0
	for k := 0; k < l; k++ {
		pR1, cR1, oR1 := e.pil[k], e.cil[k], e.oil[k]
		nR1 := float64(e.rCoeffs[k])
		if pR1[0] != 0 || cR1[0] != 0 {
			pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
			qTestTot += nR1 * QabcdBlock(
				pR1, pZb, cZb, oZb, pZb, cZb, oZb, pR1, cR1, oR1,
				pRedR1, cRedR1, oRedR1, e.prct, extZb, lZb, 1)
		}
	}

	// Compute the off-diagonal term
	numGor := l * (l - 1) / 2
	ch := make(chan float64, numGor)
	for k := 0; k < l-1; k++ {
		for j := k + 1; j < l; j++ {
			go func(k, j int) {
				pR1, cR1, oR1 := e.pil[k], e.cil[k], e.oil[k]
				pR2, cR2, oR2 := e.pil[j], e.cil[j], e.oil[j]

				if (pR1[0] == 0 && cR1[0] == 0) || (pR2[0] == 0 && cR2[0] == 0) {
					ch <- 0.0
					return
				}

				pRedR1, cRedR1, oRedR1 := ReductionBlock(pR1, cR1, oR1)
				pRedR2, cRedR2, oRedR2 := ReductionBlock(pR1, cR2, oR2)
				extR1, lR1 := ExtensionBlock(pR1, CoordsIsland(pR1, cR1, oR1))
				extR2, lR2 := ExtensionBlock(pR2, CoordsIsland(pR2, cR2, oR2))
				nR1, nR2 := float64(e.rCoeffs[k]), float64(e.rCoeffs[j])

				q := nR1*nR2*QabbaBlock(
					pR1, cR1, oR1, pR2, cR2, oR2, pRedR1, cRedR1, oRedR1, e.prct, extR2, lR2, 1) +
					nR1*nR2*QabbaBlock(
						pR2, cR2, oR2, pR1, cR1, oR1, pRedR2, cRedR2, oRedR2, e.prct, extR1, lR1, 1)
				ch <- q
			}(k, j)
		}
	}

	for i := 0; i < numGor; i++ {
		qTestTot += <-ch
	}
	close(ch)

	return qTestTot
}
