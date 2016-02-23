//
// load_molecule.go
//

package functions

import (
	"log"
	"os"
	"path"

	. "../util"
)

type MoleculeCoordinates struct {
	C, H, Br [][]float64
	All      [][]float64
}

func LoadMoleculeCoordinates(dataDir string, molecules ...string) *MoleculeCoordinates {
	l := len(molecules)
	entity := make([]string, l)
	for i, v := range molecules {
		sym := path.Join(dataDir, v)
		if f, err := os.Readlink(sym); err != nil {
			log.Fatalf("%s is not a symbolic link file.", f)
		} else {
			entity[i] = path.Join(dataDir, f)
		}
	}

	m, lm := make([][][]float64, l), 0
	Natoms = make([]int, l) // Numbers of atoms in order of appearence in Coordinates
	for i, e := range entity {
		m[i] = LoadFromCsvFile2Dim(e, ' ')
		mlen := len(m[i])
		Natoms[i] = 2 * mlen
		lm += mlen
	}

	// Concatenate all molecules data
	all := make([][]float64, 0, lm)
	for _, v := range m {
		all = append(all, v...)
	}

	return &MoleculeCoordinates{m[0], m[1], m[2], all}
}
