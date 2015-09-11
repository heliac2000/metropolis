//
// energy_pair.go
//

package functions

import (
	"math"
	"sort"

	"github.com/skelterjohn/go.matrix"

	. "../util"
)

// Compute the interaction energy of a molecule at unit cell k1, in
// character c1, in orientation o1, and a molecule at unit cel k2,
// character c2, in orientation o2. Orientations in degrees.
//
func EnergyPair(k1, k2, ch1, ch2 int, o1, o2 float64) []float64 {
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
	distIJ := Dist(coordinates)
	var minFast [][]float64
	Create2DimArray(&minFast, len(distIJ), len(distIJ[0]))
	for i := 0; i < len(distIJ); i++ {
		for j := 0; j < len(distIJ[0]); j++ {
			minFast[i][j] = Zcoulomb[i][j] / distIJ[i][j]
		}
	}

	return EigenValues(minFast)
	//speck := EigenValues(minFast)
	//fmt.Println(speck)

	// Decide if the interaction is attractive or repulsive
	// Check via support vector machines
}

// [R] dist function(3-dimension)
//
func Dist(m [][]float64) [][]float64 {
	var dist [][]float64
	l := len(m)
	Create2DimArray(&dist, l, l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				dist[i][j] = Mcut
			} else {
				dist[i][j] =
					math.Sqrt((m[i][0]-m[j][0])*(m[i][0]-m[j][0]) +
						(m[i][1]-m[j][1])*(m[i][1]-m[j][1]) +
						(m[i][2]-m[j][2])*(m[i][2]-m[j][2]))
			}
		}
	}

	return dist
}

// [R] eigen function(return eigen$values only)
//
func EigenValues(mat [][]float64) []float64 {
	l := len(mat) // Square matrix
	m := make([]float64, l*l)
	for i, k := 0, 0; i < l; i++ {
		for j := 0; j < l; j++ {
			m[k], k = mat[i][j], k+1
		}
	}

	dm := matrix.MakeDenseMatrix(m, l, l)
	_, v, err := dm.Eigen()
	if err != nil {
		return m
	}

	ev := make([]float64, l)
	for i := 0; i < l; i++ {
		ev[i] = v.Get(i, i)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(ev)))

	return ev
}
