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
	CentralPoint = 5

	// Input the hopping lattice from the unit cell

	// Number of horizontal and vertical repetitions of the unit cell to
	// generate interaction lattice (assuming unit cell with straight
	// edges)
	Nrepeat int = 100

	// Number of cuts of the lattice vectors to use to identify lattice
	// points in the lattice planes
	Nstep int = 1e5

	// Number of molecules to consider
	Nmolec int = 10

	// Number of steps in the Metropolis-Hastings algorithm
	Nmetro int = 1e4

	// Parameters for binomial random variable generation

	// How far to extend the occupied cells when computing extension set
	Npower int = 5

	// How much unit cell radius to exclude about each occupied point
	Nexclude int = 3

	// Parameters for repulsive part of interaction

	// Minimum distance that two atoms by be in before touching repulsive wall
	Mcut float64 = 1.0

	// Radius to get rid of duplicated points
	Epsilon float64 = 0.5
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
	Temp float64 = 75

	// For parallel tempering
	TempS []float64 = []float64{50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}

	// For parallel tempering
	Nparallel int = len(TempS)

	// Genereate island from UC center(use 1227 for 100 x 100 lattice, use 313 for 50 x 50 lattice)
	UCcenter int = 1227

	// Numbers of atoms in order of appearence in Coordinates
	Natoms []int = []int{56, 32, 4}

	// Assign the atomic numbers
	AtomNumber []float64 = []float64{6.0, 1.0, 35.0}

	// The numerators of the Coulomb matrices
	Zcoulomb [][]float64

	// KRLS objects
	KernelRegsRepLog Krls
	KernelRegsAtt    Krls

	// Svm object
	SvmModel Svm
)
