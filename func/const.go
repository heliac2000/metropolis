//
// Constants
//

package functions

import "runtime"

const (
	// Number of unit cells projections along a-axis # For degeneracy
	// calculation
	Nuc int = 25

	// Temperature in kelvin.
	Temp float64 = 275

	// Boltzmann ant in eV K-1
	KB float64 = 1.38e-23 / 1.60e-19

	// Scaling parameters for extension-reduction transform
	Alpha1 int = 1
	Alpha2 int = 1

	// Input the hopping lattice from the unit cell

	// Number of horizontal and vertical repetitions of the unit cell to
	// generate interaction lattice (assuming unit cell with straight
	// edges)
	Nrepeat int = 50

	// Number of cuts of the lattice vectors to use to identify lattice
	// points in the lattice planes
	Nstep int = 1e4

	// Number of molecules to consider
	Nmolec int = 10

	// Number of steps in the Metropolis-Hastings algorithm
	Nmetro int = 1e4

	// Parameters for binomial random variable generation

	// How far to extend the occupied cells when computing extension set
	Npower int = 15
)

// Global variables
//
var (
	// Number of Concurrency
	NumConcurrency = runtime.NumCPU()
)
