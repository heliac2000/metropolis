//
// canonical_order.go
//

package functions

import (
	"sort"

	. "../util"
)

// Arrange the elements of C1 in order of decreasing size, append one
// element of size 0. C1 should be in form of a position list,
// character list, and orientation list.
//
func CanonicalOrder(ctemp, chtemp [][]int, otemp [][]float64) ([][]int, [][]int, [][]float64) {
	l := len(ctemp)
	clengths := make([]int, l)
	for k := 0; k < l; k++ {
		if len(ctemp[k]) == 1 && ctemp[k][0] == 0 {
			clengths[k] = 0
		} else {
			clengths[k] = len(ctemp[k])
		}
	}

	csort_map := make(map[int][]int)
	for i := 0; i < l; i++ {
		csort_map[clengths[i]] = append(csort_map[clengths[i]], i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(clengths)))

	csortIND, pre := make([]int, 0, l), clengths[0]
	csortIND = append(csortIND, csort_map[pre]...)
	for _, v := range clengths {
		if v != pre {
			csortIND = append(csortIND, csort_map[v]...)
			pre = v
		}
	}

	csort, chsort, osort :=
		make([][]int, 0, l+1), make([][]int, 0, l+1), make([][]float64, 0, l+1)

	for k := 0; k < l; k++ {
		csort = append(csort, CopyVector(ctemp[csortIND[k]]).([]int))
		chsort = append(chsort, CopyVector(chtemp[csortIND[k]]).([]int))
		osort = append(osort, CopyVector(otemp[csortIND[k]]).([]float64))
		if len(csort[k]) == 1 && csort[k][0] == 0 {
			goto exit
		}
	}

	csort = append(csort, []int{0})
	chsort = append(chsort, []int{0})
	osort = append(osort, []float64{0})

exit:
	return csort, chsort, osort
}
