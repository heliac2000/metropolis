//
// canonical.go
//

package functions

// Canonical type
//
type Canonical struct {
	Pos, Chr [][]int
	Ori      [][]float64
}

// Methods
//

// Constructor
func NewCanonical(pos, chr [][]int, ori [][]float64) Canonical {
	return Canonical{Pos: pos, Chr: chr, Ori: ori}
}

// Clone
func (src Canonical) Dup() Canonical {
	r := len(src.Pos)
	pos, chr, ori := make([][]int, r), make([][]int, r), make([][]float64, r)
	for i := 0; i < r; i++ {
		c := len(src.Pos[i])
		pos[i], chr[i], ori[i] = make([]int, c), make([]int, c), make([]float64, c)
		copy(pos[i], src.Pos[i])
		copy(chr[i], src.Chr[i])
		copy(ori[i], src.Ori[i])
	}

	return Canonical{Pos: pos, Chr: chr, Ori: ori}
}

// Return all items
func (c Canonical) Explode() ([][]int, [][]int, [][]float64) {
	return c.Pos, c.Chr, c.Ori
}

// Pack to canonical object
func CanonicalImplode(pos, chr [][]int, ori [][]float64, args ...interface{}) (Canonical, []interface{}) {
	return Canonical{Pos: pos, Chr: chr, Ori: ori}, args
}
