//
// load_molecule.go
//

package functions

import . "../util"

type MoleculeCoordinates struct {
	C, H, Br [][]float64
}

func LoadMoleculeCoordinates(ccarts, hcarts, brcarts string) *MoleculeCoordinates {
	C := LoadFromCsvFile2Dim(ccarts, ' ')
	H := LoadFromCsvFile2Dim(hcarts, ' ')
	Br := LoadFromCsvFile2Dim(brcarts, ' ')

	return &MoleculeCoordinates{C, H, Br}
}
