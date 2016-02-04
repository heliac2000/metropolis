//
// load_molecule.go
//

package functions

import (
	"path"

	. "../util"
)

type MoleculeCoordinates struct {
	C, H, Br [][]float64
	All      [][]float64
}

func LoadMoleculeCoordinates(dataDir, cCarts, hCarts, brCarts string) *MoleculeCoordinates {
	C := LoadFromCsvFile2Dim(path.Join(dataDir, cCarts), ' ')
	H := LoadFromCsvFile2Dim(path.Join(dataDir, hCarts), ' ')
	Br := LoadFromCsvFile2Dim(path.Join(dataDir, brCarts), ' ')

	all := make([][]float64, 0, len(C)+len(H)+len(Br))
	all = append(append(append(all, C...), H...), Br...)

	return &MoleculeCoordinates{C, H, Br, all}
}
