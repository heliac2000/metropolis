//
// sort_index_return.go
//

package util

import "sort"

// R's `sort(x, index.return = TRUE)$ix'
//
func SortIndexReturn(y []int) []int {
	x := CopyVectorInt(y)
	sorted, l := make(map[int][]int), len(x)
	for i := 0; i < l; i++ {
		sorted[x[i]] = append(sorted[x[i]], i)
	}
	sort.Sort(sort.IntSlice(x))

	idx, pre := make([]int, 0, l), x[0]
	idx = append(idx, sorted[pre]...)
	for _, v := range x {
		if v != pre {
			idx = append(idx, sorted[v]...)
			pre = v
		}
	}

	return idx
}
