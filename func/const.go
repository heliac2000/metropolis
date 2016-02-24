//
// Constants
//

package functions

import (
	"math/rand"
	"time"
)

const (
	// Number of unit cells projections along a-axis # For degeneracy
	// calculation
	Nuc int = 50

	// Boltzmann ant in eV K-1
	KB float64 = 1.38e-23 / 1.60e-19

	// Scaling parameters for extension-reduction transform
	Alpha1 float64 = 1
	Alpha2 float64 = 1

	// Character of unit cell central point
	CentralPoint int = 5

	// Input the hopping lattice from the unit cell

	// Number of horizontal and vertical repetitions of the unit cell to
	// generate interaction lattice (assuming unit cell with straight
	// edges)
	Nrepeat int = 100

	// Number of cuts of the lattice vectors to use to identify lattice
	// points in the lattice planes
	//Nstep int = 1e5

	// Number of molecules to consider
	Nmolec int = 10

	// Number of steps in the Metropolis-Hastings algorithm
	Nmetro int = 1e4

	// Parameters for binomial random variable generation

	// How far to extend the occupied cells when computing extension set
	Npower int = 5

	// Genereate island from UC center(use 1227 for 100 x 100 lattice, use 313 for 50 x 50 lattice)
	UCcenter int = 1227

	// How much unit cell radius to exclude about each occupied point
	//Nexclude int = 3

	// Parameters for repulsive part of interaction

	// Minimum distance that two atoms by be in before touching repulsive wall
	Mcut float64 = 1.0

	// Maximum distance allowable between two atoms ('cut-off distance')
	McomCUT float64 = 8.0

	// Radius to get rid of duplicated points
	Epsilon float64 = 0.5

	// Set to TRUE if wish to use low-dimensional representation based on PCA
	PcaRep bool = true

	// Energy for unstable region
	Eunstable float64 = 10

	// Detection limit for angles
	Detect float64 = 0.01
)

// Data files
//
const (
	// UnitCell 2
	UnitCell2File = "UnitCell2.csv"

	// Lattice Vectors
	LatticeVectorsFile = "PrecursorUnitCellAxes.csv"
	LatticeFile        = "Lattice.csv"

	// Character of unit cell
	CharacterFile = "Character.csv"

	// Sequence of power matrices
	AdjCumlFile = "AdjCuml.csv"

	// Coordinates of atoms in molecule
	CCoords  = "Molecule_01.csv" // Typically, Carbon
	HCoords  = "Molecule_02.csv" // Typically, Hydrogen
	BrCoords = "Molecule_03.csv" // Typically, Bromine

	// KRLS objects
	KernelRegsRepLogFile = "KernelregSRepLog.json"
	KernelRegsAttFile    = "KernelregSAtt.json"

	// SVM objects
	SvmModelOpFile  = "SvmModelOp.json"
	SvmModelNzpFile = "SvmModelNzp.json"
	SvmModelUspFile = "SvmModelUsp.json"

	// PrComp objects
	XeigPcFile = "Xeigpc.json"
)

// Atom identities
const (
	AtomC int = iota
	AtomH
	AtomBr
)

// Global variables
//
var (
	// Number of Concurrency
	//NumConcurrency int = runtime.NumCPU()

	// Random seed
	Rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Input data
	Inp *InitData

	// Temperature in kelvin.
	//Temp float64 = 75

	// For parallel tempering
	TempS []float64

	// For parallel tempering
	Nparallel int

	// Numbers of atoms in order of appearence in Coordinates
	Natoms []int

	// Assign the atomic numbers
	AtomNumber []float64

	// The numerators of the Coulomb matrices
	Zcoulomb [][]float64

	// KRLS objects
	KernelRegsRepLog, KernelRegsAtt Krls

	// Svm object
	SvmModel, SvmModelOp, SvmModelNzp, SvmModelUsp Svm

	// Number of principal components to use in the force field
	Npcs int

	// PrComp object
	XeigPc PrComp

	// List of variables
	Vars [][]int

	// Atomic Numbers
	AtomicNumbers map[string]float64 = map[string]float64{
		"H": 1, "He": 2, "Li": 3, "Be": 4, "B": 5, "C": 6, "N": 7, "O": 8, "F": 9, "Ne": 10,
		"Na": 11, "Mg": 12, "Al": 13, "Si": 14, "P": 15, "S": 16, "Cl": 17, "Ar": 18, "K": 19, "Ca": 20,
		"Sc": 21, "Ti": 22, "V": 23, "Cr": 24, "Mn": 25, "Fe": 26, "Co": 27, "Ni": 28, "Cu": 29, "Zn": 30,
		"Ga": 31, "Ge": 32, "As": 33, "Se": 34, "Br": 35, "Kr": 36, "Rb": 37, "Sr": 38, "Y": 39, "Zr": 40,
		"Nb": 41, "Mo": 42, "Tc": 43, "Ru": 44, "Rh": 45, "Pd": 46, "Ag": 47, "Cd": 48, "In": 49, "Sn": 50,
		"Sb": 51, "Te": 52, "I": 53, "Xe": 54, "Cs": 55, "Ba": 56, "La": 57, "Ce": 58, "Pr": 59, "Nd": 60,
		"Pm": 61, "Sm": 62, "Eu": 63, "Gd": 64, "Tb": 65, "Dy": 66, "Ho": 67, "Er": 68, "Tm": 69, "Yb": 70,
		"Lu": 71, "Hf": 72, "Ta": 73, "W": 74, "Re": 75, "Os": 76, "Ir": 77, "Pt": 78, "Au": 79, "Hg": 80,
		"Tl": 81, "Pb": 82, "Bi": 83, "Po": 84, "At": 85, "Rn": 86, "Fr": 87, "Ra": 88, "Ac": 89, "Th": 90,
		"Pa": 91, "U": 92, "Np": 93, "Pu": 94, "Am": 95, "Cm": 96, "Bk": 97, "Cf": 98, "Es": 99, "Fm": 100,
		"Md": 101, "No": 102, "Lr": 103, "Rf": 104, "Db": 105, "Sg": 106, "Bh": 107, "Hs": 108, "Mt": 109, "Ds": 110,
		"Rg": 111, "Cp": 112, "Uut": 113, "Uuq": 114, "Uup": 115, "Uuh": 116, "Uus": 117, "Uuo": 118,
	}
)
