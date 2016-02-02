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
	return Canonical{
		pos: Copy2DimArrayInt(src.pos),
		chr: Copy2DimArrayInt(src.chr),
		ori: Copy2DimArrayFloat(src.ori),
	}
}

// Return all items
func (c Canonical) Explode() ([][]int, [][]int, [][]float64) {
	return c.pos, c.chr, c.ori
}

// Pack to canonical object
func CanonicalImplode(pos, chr [][]int, ori [][]float64, args ...interface{}) (Canonical, []interface{}) {
	return Canonical{pos: pos, chr: chr, ori: ori}, args
}
