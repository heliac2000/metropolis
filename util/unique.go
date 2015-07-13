package util

// Unique
//
func Unique(arr []int) []int {
	uniq := make([]int, 0, len(arr))
	for _, v := range arr {
		if !Member(v, uniq) {
			uniq = append(uniq, v)
		}
	}
	return uniq
}
