//
// energy_island_pca.go
//

package functions

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math"
)

// Calculate energy of an island (Island[[1]] = positions, Island[[2]]
// = characters, Island[[2]] = orientations). Only use this if PCA_REP
// = TRUE
//
func EnergyIslandPCA(pos, chr []int, ori []float64) float64 {
	// Adsorption energy
	ene := 0.0
	for k := 0; k < len(pos); k++ {
		// Incase new UnitCell2 file format is used
		for _, u := range Inp.UnitCell2 {
			// u[0]: color, u[3]: angle, u[4]: energy
			if int(u[0]) == Inp.Character[chr[k]]+1 && math.Abs(u[3]-ori[k]) < Detect {
				ene += u[4]
			}
		}
	}

	if len(pos) <= 1 {
		return ene
	}

	// Interaction energy
	for k := 0; k < len(pos)-1; k++ {
		for j := k + 1; j < len(pos); j++ {
			h := getHashEnergyPair(pos[k], pos[j], chr[k], chr[j], ori[k], ori[j])
			v, ok := 0.0, false
			if v, ok = energyPairReduceMap[h]; !ok {
				v = EnergyPairReduce(pos[k], pos[j], chr[k], chr[j], ori[k], ori[j])
				energyPairReduceMap[h] = v
			}
			ene += v
		}
	}

	return ene
}

//
// Memorize
//
var energyPairReduceMap map[string]float64 = make(map[string]float64)

func getHashEnergyPair(pos1, pos2, chr1, chr2 int, ori1, ori2 float64) string {
	b := make([]byte, 48)

	for i, v := range []int{pos1, pos2, chr1, chr2} {
		binary.LittleEndian.PutUint64(b[(i*8):], uint64(v))
	}
	for i, v := range []float64{ori1, ori2} {
		binary.LittleEndian.PutUint64(b[((i+4)*8):], math.Float64bits(v))
	}

	return fmt.Sprintf("%x", md5.Sum(b))
}

// func getHash(pos, chr []int, ori []float64) string {
// 	b := make([]byte, (len(pos)+len(chr)+len(ori))*8)
//
// 	for i, v := range pos {
// 		binary.LittleEndian.PutUint64(b[(i*8):], uint64(v))
// 	}
// 	for i, v := range chr {
// 		binary.LittleEndian.PutUint64(b[((i+len(pos))*8):], uint64(v))
// 	}
// 	for i, v := range ori {
// 		binary.LittleEndian.PutUint64(b[((i+len(pos)+len(chr))*8):], math.Float64bits(v))
// 	}
//
// 	return fmt.Sprintf("%x", md5.Sum(b))
// }
