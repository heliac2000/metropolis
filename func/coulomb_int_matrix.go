//
// coulomb_int_matrix.go
//

package functions

import (
	"math"

	. "../util"
)

// Input the functions for the PCA case. Must be extremely careful
// with the ordering of the atoms in the molecules!
//
// Compute the interaction energy of a molecule at unit cell k1, in
// character c1, in orientation o1, and a molecule at unit cel k2,
// character c2, in orientation o2. Orientations in degrees.
//
// Generate the interaction Coulomb matrix for the pair.
//
func CoulombIntMatrix(k1, k2, ch1, ch2 int, o1, o2 float64) ([][]float64, bool) {
	deltaxy1 := []float64{
		Inp.UnitCell[CentralPoint][1] - Inp.UnitCell[ch1][1],
		Inp.UnitCell[CentralPoint][2] - Inp.UnitCell[ch1][2],
	}
	deltaxy2 := []float64{
		Inp.UnitCell[CentralPoint][1] - Inp.UnitCell[ch2][1],
		Inp.UnitCell[CentralPoint][2] - Inp.UnitCell[ch2][2],
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
	coordinates := Create2DimArrayFloat(l, 3)

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
	// Min_fast is the standard Coulomb matirx. Now, convert to
	// interaction Coulomb matrix.
	minFast := Create2DimArrayFloat(len(distIJ), len(distIJ[0]))
	for i := 0; i < len(distIJ); i++ {
		for j := 0; j < len(distIJ[0]); j++ {
			minFast[i][j] = Zcoulomb[i][j] / distIJ[i][j]
		}
	}

	// Molecule 1 atoms, Molecule 2 atoms
	cm := Create2DimArrayFloat(len(m1), len(m1))
	region := append(SeqInt(0, lc-1, 1),
		append(SeqInt(2*lc, 2*lc+lh-1, 1), SeqInt(2*(lc+lh), 2*(lc+lh)+lbr-1, 1)...)...)
	for i, idx := range region {
		copy(cm[i][0:], minFast[idx][lc:(2*lc)])
		copy(cm[i][lc:], minFast[idx][(2*lc+lh):(2*(lc+lh))])
		copy(cm[i][(lc+lh):], minFast[idx][(2*(lc+lh)+lbr):])
	}

	coords := CoordsGen(k1, k2, ch1, ch2, o1, o2)
	dist12, _ := GetKnnx(coords[:len(m1)], coords[len(m1):], 1)

	return cm, (MinFloat(Transpose(dist12)[0]...) > McomCUT)
}
