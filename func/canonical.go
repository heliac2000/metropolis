//
// canonical.go
//

package functions

import . "../util"

// Canonical type
//
type Canonical struct {
	pos, chr [][]int
	ori      [][]float64
}

// Methods
//

// Constructor
func NewCanonical(pos, chr [][]int, ori [][]float64) Canonical {
	return Canonical{pos: pos, chr: chr, ori: ori}
}

// Clone
func (src Canonical) Dup() Canonical {
	var dst Canonical
	Copy2DimArray(&dst.pos, src.pos)
	Copy2DimArray(&dst.chr, src.chr)
	Copy2DimArray(&dst.ori, src.ori)

	return dst
}

// Return all items
func (c Canonical) Explode() ([][]int, [][]int, [][]float64) {
	return c.pos, c.chr, c.ori
}

// Pack to canonical object
func CanonicalImplode(pos, chr [][]int, ori [][]float64, args ...interface{}) (Canonical, []interface{}) {
	return Canonical{pos: pos, chr: chr, ori: ori}, args
}
