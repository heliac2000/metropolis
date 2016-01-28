//
// expand_grid.go
//

package util

func ExpandGrid(arr [][]int) [][]int {
	r, c := 1, len(arr)
	for i := 0; i < c; i++ {
		r *= len(arr[i])
	}

	m := Create2DimArrayInt(r, c)
	tick := 1
	for i := 0; i < c; i++ {
		for j := 0; j < r/tick; j++ {
			for k := 0; k < tick; k++ {
				m[tick*j+k][i] = arr[i][j%len(arr[i])]
			}
		}
		tick *= len(arr[i])
	}

	return m
}

func Flatten(arr [][]int) []int {
	r, c := len(arr), len(arr[0])
	vec := make([]int, 0, r*c)

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			vec = append(vec, arr[j][i])
		}
	}

	return vec
}
