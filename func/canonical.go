//
// canonical.go
//

package functions

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
	r := len(src.pos)
	pos, chr, ori := make([][]int, r), make([][]int, r), make([][]float64, r)
	for i := 0; i < r; i++ {
		c := len(src.pos[i])
		pos[i], chr[i], ori[i] = make([]int, c), make([]int, c), make([]float64, c)
		copy(pos[i], src.pos[i])
		copy(chr[i], src.chr[i])
		copy(ori[i], src.ori[i])
	}

	return Canonical{pos: pos, chr: chr, ori: ori}
}

// Return all items
func (c Canonical) Explode() ([][]int, [][]int, [][]float64) {
	return c.pos, c.chr, c.ori
}

// Pack to canonical object
func CanonicalImplode(pos, chr [][]int, ori [][]float64, args ...interface{}) (Canonical, []interface{}) {
	return Canonical{pos: pos, chr: chr, ori: ori}, args
}
