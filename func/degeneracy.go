//
// degeneracy.go
//

package functions

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math"

	. "../util"
)

// Compute the degeneracy factor (from equation 63, page 284 of notes)
//
func Degeneracy(pos, chr [][]int, ori [][]float64) float64 {
	h := getHashDegeneracy(pos, chr, ori)
	if v, ok := DegeneracyMap[h]; ok {
		return v
	}

	// Determine the number of non-zero islands in Ctest. Important for
	// Ctest to be ordered.
	clength := 0
	for k := 0; k < len(pos); k++ {
		sum := 0
		for _, v := range pos[k] {
			sum += v
		}
		if sum == 0 {
			clength = k
			break
		}
	}

	// Remove zeroes from canonical representation.
	// rot: Store all rotated islands, rotInd: Unique indices
	rotInd, rot := make([][]int, clength), make([][]int, clength)
	for k := 0; k < clength; k++ {
		rotInd[k], rot[k] = UniqueOrientations(pos[k], chr[k], ori[k])
	}

	// Now, determine all unique combinations of indices from Rot
	rotCombs, degF := ExpandGrid(rotInd), 0.0
	for _, indsk := range rotCombs {
		ppos := make([][]int, clength)
		for j := 0; j < clength; j++ {
			if ind := rotInd[j][indsk[j]]; ind == 0 {
				ppos[j] = pos[j]
			} else {
				ppos[j] = rot[j]
			}
		}
		factors := TransDen(ppos, chr[:clength], ori[:clength])
		addF := 1.0
		for _, v := range factors {
			addF *= Factorial(v)
		}
		degF += 1.0 / addF
	}

	// Degeneracy factor
	v := Factorial(clength) * ApproXnCr(Nuc*Nuc, clength) * degF
	DegeneracyMap[h] = v
	return v
}

//
// Memorize
//
var DegeneracyMap map[string]float64 = make(map[string]float64)

func getHashDegeneracy(pos, chr [][]int, ori [][]float64) string {
	h := make([]byte, 0, 3*8*(len(pos)*len(pos[0])))
	for i := 0; i < len(pos); i++ {
		l := len(pos[i])
		p, c, o := make([]byte, l*8), make([]byte, l*8), make([]byte, l*8)
		for j := 0; j < l; j++ {
			binary.LittleEndian.PutUint64(p[(j*8):], uint64(pos[i][j]))
			binary.LittleEndian.PutUint64(c[(j*8):], uint64(chr[i][j]))
			binary.LittleEndian.PutUint64(o[(j*8):], math.Float64bits(ori[i][j]))
		}
		h = append(append(append(h, p...), c...), o...)
	}

	return fmt.Sprintf("%x", md5.Sum(h))
}
