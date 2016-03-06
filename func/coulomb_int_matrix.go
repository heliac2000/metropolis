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

	for i, l := 0, 0; i < len(Inp.MoleculeCoordinates.Mol); i++ {
		mlen := len(Inp.MoleculeCoordinates.Mol[i])
		copy(coordinates[(2*l):], m1[l:(l+mlen)])
		copy(coordinates[((2*l)+mlen):], m2[l:(l+mlen)])
		l += mlen
	}

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
	region, l := make([]int, 0, 2*len(m1)), 0
	for _, v := range Inp.MoleculeCoordinates.Mol {
		mlen := len(v)
		region = append(region, SeqInt(2*l, 2*l+mlen-1, 1)...)
		l += mlen
	}

	for i, idx := range region {
		l := 0
		for _, v := range Inp.MoleculeCoordinates.Mol {
			mlen := len(v)
			copy(cm[i][l:], minFast[idx][(2*l+mlen):(2*(l+mlen))])
			l += mlen
		}
	}

	coords := CoordsGen(k1, k2, ch1, ch2, o1, o2)
	dist12, _ := GetKnnx(coords[:len(m1)], coords[len(m1):], 1)

	return cm, (MinFloat(Transpose(dist12)[0]...) > McomCUT)
}
