//
// energy_pair.go
//

package functions

import (
	"math"

	. "../util"
)

// Compute the interaction energy of a molecule at unit cell k1, in
// character c1, in orientation o1, and a molecule at unit cel k2,
// character c2, in orientation o2. Orientations in degrees.
//
func EnergyPair(k1, k2, ch1, ch2 int, o1, o2 float64) float64 {
	// First, need to get the coordinates of the molecules
	deltaxy1 := []float64{
		Inp.UnitCell2[3][1] - Inp.UnitCell2[ch1][1],
		Inp.UnitCell2[3][2] - Inp.UnitCell2[ch1][2],
	}
	deltaxy2 := []float64{
		Inp.UnitCell2[3][1] - Inp.UnitCell2[ch2][1],
		Inp.UnitCell2[3][2] - Inp.UnitCell2[ch2][2],
	}

	coords1 := []float64{
		Inp.UnitCellCoords[k1][0] - deltaxy1[0],
		Inp.UnitCellCoords[k1][1] - deltaxy1[1],
	}
	coords2 := []float64{
		Inp.UnitCellCoords[k2][0] - deltaxy2[0],
		Inp.UnitCellCoords[k2][1] - deltaxy2[1],
	}

	m1 := ShiftMCpos(Inp.MoleculeCoordinates, coords1)
	m2 := ShiftMCpos(Inp.MoleculeCoordinates, coords2)
	m1 = RotateZ(m1, o1*math.Pi/180.0)
	m2 = RotateZ(m2, o2*math.Pi/180.0)

	// Put together the cooridinate matrix
	l := len(m1) * 2 // l = (lc + lh + lbr) * 2
	var coordinates [][]float64
	Create2DimArray(&coordinates, l, 3)

	lc := len(Inp.MoleculeCoordinates.C)
	copy(coordinates, m1[:lc])
	copy(coordinates[lc:], m2[:lc])

	lh := len(Inp.MoleculeCoordinates.H)
	copy(coordinates[(2*lc):], m1[lc:(lc+lh)])
	copy(coordinates[(2*lc+lh):], m2[lc:(lc+lh)])

	lbr := len(Inp.MoleculeCoordinates.Br)
	copy(coordinates[(2*(lc+lh)):], m1[(lc+lh):])
	copy(coordinates[(2*(lc+lh)+lbr):], m2[(lc+lh):])

	// Convert coordinate matrix into a Coulomb matrix
	distIJ := Dist(coordinates, Mcut)
	var minFast [][]float64
	Create2DimArray(&minFast, len(distIJ), len(distIJ[0]))
	for i := 0; i < len(distIJ); i++ {
		for j := 0; j < len(distIJ[0]); j++ {
			minFast[i][j] = Zcoulomb[i][j] / distIJ[i][j]
		}
	}

	speck := [][]float64{EigenValues(minFast)}

	// Decide if the interaction is attractive or repulsive
	// Check via support vector machines
	// int_type = predict(svm_model, newdata = speck)
	intType := "attractive"

	// For repulsive type, predict log of interaction energy
	eint := 0.0
	if intType == "repulsive" {
		eint = math.Exp(KernelRegsRepLog.Predict(speck))
	} else if intType == "attractive" {
		eint = KernelRegsAtt.Predict(speck)
	}

	return eint
}
