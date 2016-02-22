//
// load_molecule.go
//

package functions

import (
	"fmt"
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

	// Numbers of atoms in order of appearence in Coordinates
	Natoms = []int{len(C) * 2, len(H) * 2, len(Br) * 2}
	fmt.Println(Natoms)

	return &MoleculeCoordinates{C, H, Br, all}
}
