//
// load_molecule.go
//

package functions

import . "../util"

type MoleculeCoordinates struct {
	C, H, Br [][]float64
	All      [][]float64
}

func LoadMoleculeCoordinates(ccarts, hcarts, brcarts string) *MoleculeCoordinates {
	C := LoadFromCsvFile2Dim(ccarts, ' ')
	H := LoadFromCsvFile2Dim(hcarts, ' ')
	Br := LoadFromCsvFile2Dim(brcarts, ' ')

	all := make([][]float64, 0, len(C)+len(H)+len(Br))
	all = append(append(append(all, C...), H...), Br...)

	return &MoleculeCoordinates{C, H, Br, all}
}
