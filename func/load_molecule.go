//
// load_molecule.go
//

package functions

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"

	. "../util"
)

type MoleculeCoordinates struct {
	Mol [][][]float64
	All [][]float64
}

func LoadMoleculeCoordinates(dataDir string) *MoleculeCoordinates {
	entity := make([]string, 0)
	Atoms, AtomNumber = make([]string, 0), make([]float64, 0)
	r := regexp.MustCompile(`^(.+)coords.*\.csv$`)
	nMol := 1
	for {
		molFile := path.Join(dataDir, fmt.Sprintf("Molecule_%02d.csv", nMol))
		if _, err := os.Stat(molFile); err != nil {
			nMol--
			break
		}
		if f, err := os.Readlink(molFile); err != nil {
			log.Fatalf("%s is not a symbolic link file.", molFile)
		} else {
			atom := string(r.FindSubmatch([]byte(f))[1])
			an, ok := AtomicNumbers[atom]
			if !ok {
				log.Fatalf("Illegal atom %s.", atom)
			}
			Atoms = append(Atoms, atom)
			AtomNumber = append(AtomNumber, an)
			entity = append(entity, path.Join(dataDir, f))
		}
		nMol++
	}

	if nMol < 2 {
		log.Fatalln("Number of Molecule must be greater than 1.")
	}

	m, lm := make([][][]float64, nMol), 0
	Natoms = make([]int, nMol) // Numbers of atoms in order of appearence in Coordinates
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

	return &MoleculeCoordinates{m, all}
}
