//
// reaction_class_id.go
//

package functions

import (
	"reflect"
	"sort"
)

// Next part: Calculation transistion probability based on reaction class
// Key:
const (
	REACT_UNKNOWN int = iota - 1
	REACT_CLASS1
	REACT_CLASS2
	REACT_CLASS3
	REACT_CLASS4
	REACT_CLASS5
	REACT_CLASS6
	REACT_CLASS7
	REACT_CLASS8
	REACT_CLASS9
)

var classKey = [][]int{
	{-1, -1, 1, 1},
	{-1, -1, 2},
	{-1, -1, 1},
	{-2, 1, 1},
	{-2, 1},
	{-1, 1, 1},
	{-1, 2},
	{-1, 1},
}

// Identify reaction class according to scheme above
//
//
func ReactionClassID(diff []int) int {
	// Fetch the reaction difference vector
	//_, _, _, _, _, diff := ReactionRepresentation(prct, crct, orct, ppdt, cpdt, opdt)

	// Later, use the full data in dIJ to calculate the probabilities
	dijp := make([]int, 0, len(diff))
	for _, v := range diff {
		if v != 0 {
			dijp = append(dijp, v)
		}
	}

	l := len(dijp)
	if l == 0 {
		return REACT_CLASS9
	} else if l > 4 || l == 1 {
		return REACT_UNKNOWN
	}

	sort.Ints(dijp)
	for i, v := range classKey {
		if len(v) == l && reflect.DeepEqual(dijp, v) {
			return i
		}
	}

	return REACT_UNKNOWN
}
