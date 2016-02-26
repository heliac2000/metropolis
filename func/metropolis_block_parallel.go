//
// metropolis_block_parallel.go
//

package functions

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	. "../util"
)

// Do Metropolis-Hastings sampling with 1-step shifts (slow, uses
// exact PF; uses parallel tempering)
//
func MetropolisBlockParallel(N int, eout, cout string) {
	// random seed
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Truncate and open files
	fn := getFuncName()
	truncateFile(cout, fn)
	truncateFile(eout, fn)
	fpCout, fpEout := openFiles(cout, eout)
	coutWriter := csv.NewWriter(fpCout)
	defer fpCout.Close()
	defer fpEout.Close()

	tick := 1000
	if N < tick {
		tick = N
	}

	coutP, eoutP := make([][]Canonical, Nparallel), make([][]float64, Nparallel)
	for k := 0; k < Nparallel; k++ {
		coutP[k] = make([]Canonical, tick)
		eoutP[k] = make([]float64, tick)
		coutP[k][0] = NewCanonical(CanonicalOrder(CanonicalGen()))
		eoutP[k][0] = EnergyCanonical(coutP[k][0].Explode())
	}

	for n := 0; n < N/tick; n++ {
		for k := 1; k < tick; k++ {
			if k%100 == 99 {
				log.Printf("n = %4d/N = %5d\n", n*tick+k+1, N)
			}

			// Draw a uniform random variable to decide to do the tempering
			if randFloat64(rnd) < 0.5 {
				// Do an ordinary extension of a uniformly chosen element
				schoose := rnd.Intn(Nparallel)
				cout := make([]Canonical, len(coutP[schoose]))
				for i := 0; i < k; i++ {
					cout[i] = coutP[schoose][i].Dup()
				}
				eout := CopyVectorFloat(eoutP[schoose])
				for h := 0; h < Nparallel; h++ {
					if h != schoose {
						coutP[h][k] = coutP[h][k-1].Dup()
						eoutP[h][k] = eoutP[h][k-1]
					}
				}

				pos, chr, ori, _, _ := ExtensionReductionBlock(cout[k-1].Explode())
				coutTemp, _ := CanonicalImplode(CanonicalOrder(pos, chr, ori))
				qC1C2 := ExtensionReductionProbabilityReaction(cout[k-1].Pos, cout[k-1].Chr, cout[k-1].Ori, coutTemp.Pos, coutTemp.Chr, coutTemp.Ori)
				qC2C1 := ExtensionReductionProbabilityReaction(coutTemp.Pos, coutTemp.Chr, coutTemp.Ori, cout[k-1].Pos, cout[k-1].Chr, cout[k-1].Ori)

				ax, ay := eout[k-1], EnergyCanonical(coutTemp.Explode())

				lctemp := 0
				for r := 0; r < len(coutTemp.Pos); r++ {
					if len(coutTemp.Pos[r]) > 1 || coutTemp.Pos[r][0] != 0 {
						lctemp++
					}
				}

				lcprev := 0
				for r := 0; r < len(cout[k-1].Pos); r++ {
					if len(cout[k-1].Pos[r]) > 1 || cout[k-1].Pos[r][0] != 0 {
						lcprev++
					}
				}

				// New invisible islands approximations
				pix := Degeneracy(cout[k-1].Explode())
				piy := Degeneracy(coutTemp.Explode())

				alpha := piy * qC2C1 * math.Exp(-(ay-ax)/(KB*TempS[schoose])) / (pix * qC1C2)
				if alpha > 1.0 {
					alpha = 1.0
				}

				if randFloat64(rnd) < alpha {
					coutP[schoose][k] = coutTemp
					eoutP[schoose][k] = ay
				} else {
					// Incase of rejection
					coutP[schoose][k] = cout[k-1].Dup()
					eoutP[schoose][k] = ax
				}
			} else {
				// Do the swap update
				s1, s2 := rnd.Intn(Nparallel), 0
				temp1 := TempS[s1]
				switch {
				case s1 == 0:
					s2 = 1
				case s1 == (Nparallel - 1):
					s2 = Nparallel - 2
				default:
					rmn := []int{s1 - 1, s1 + 1}
					s2 = rmn[rnd.Intn(len(rmn))]
				}

				temp2 := TempS[s2]
				// Get the new states
				c1, c2 := coutP[s1][k-1], coutP[s2][k-1]
				// Get the lengths
				l1, l2 := 0, 0
				for h := 0; h < len(c1.Pos); h++ {
					if len(c1.Pos[h]) > 1 || c1.Pos[h][0] != 0 {
						l1++
					}
				}
				for h := 0; h < len(c2.Pos); h++ {
					if len(c2.Pos[h]) > 1 || c2.Pos[h][0] != 0 {
						l2++
					}
				}

				// Energy of state 1 with temperature 1
				// Energy of state 2 with temperature 2
				e11, e22 := eoutP[s1][k-1], eoutP[s2][k-1]

				// Energy of state 2 with temperature 1
				// Energy of state 1 with temperature 2
				e12 := EnergyCanonical(c1.Explode())
				e21 := EnergyCanonical(c2.Explode())

				// [R] r12 = h1x2*h2x1/(h1x1*h2x2) * exp(-(E21 - E11)/(kB*Temp1) - (E12 - E22)/(kB*Temp2))
				// Problems here with infinities!
				r12 := math.Exp(-(e21-e11)/(KB*temp1) - (e12-e22)/(KB*temp2))
				alpha := 1.0
				if r12 < alpha {
					alpha = r12
				}

				for h := 0; h < Nparallel; h++ {
					coutP[h][k] = coutP[h][k-1].Dup()
					eoutP[h][k] = eoutP[h][k-1]
				}

				if randFloat64(rnd) < alpha {
					coutP[s1][k], coutP[s2][k] = c2.Dup(), c1.Dup()
					eoutP[s1][k], eoutP[s2][k] = e21, e12
				}
			}
		}

		// Save data
		for i := 0; i < Nparallel; i++ {
			appendArrayToFile(eoutP[i], fpEout)
			appendListWithCSV(coutP[i], coutWriter)
		}

		// Reset
		for i := 0; i < Nparallel; i++ {
			coutP[i][0], eoutP[i][0] = coutP[i][tick-1].Dup(), eoutP[i][tick-1]
			for j := 1; j < tick; j++ {
				coutP[i][j], eoutP[i][j] = Canonical{}, 0
			}
		}
	}
}

// Golang's rand.Float64() method returns a pseudo-random number in [0.0, 1.0).
// randFloat64 function returns a pseudo-random number in (0.0, 1.0), same as R's
// runif() function.
//
func randFloat64(r *rand.Rand) float64 {
	for {
		if f := float64(r.Int63()) / (1 << 63); f < 1.0 {
			return f
		}
	}
}

func getFuncName() string {
	pc := make([]uintptr, 2)
	runtime.Callers(2, pc)
	fname := runtime.FuncForPC(pc[0]).Name()
	strs := strings.Split(fname, ".")
	if l := len(strs); l > 0 {
		fname = strs[l-1]
	}

	return fname
}

func truncateFile(f, fn string) {
	if FileExists(f) {
		if fi, err := os.Stat(f); err == nil && fi.Mode().IsRegular() {
			if err := os.Truncate(f, 0); err != nil {
				log.Fatalf("[%s] %s", fn, err.Error())
			}
		}
	}
}

// Save data
//
func openFiles(cout, eout string) (*os.File, *os.File) {
	fpCout, err := os.OpenFile(cout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("[%s] %s", getFuncName(), err.Error())
	}

	fpEout, err := os.OpenFile(eout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("[%s] %s", getFuncName(), err.Error())
	}

	return fpCout, fpEout
}

func appendArrayToFile(arr []float64, fpEout *os.File) {
	for _, v := range arr {
		fmt.Fprintf(fpEout, "%.22f\n", v)
	}
}

func appendListWithCSV(can []Canonical, writer *csv.Writer) {
	for i := 0; i < len(can); i++ {
		r := len(can[i].Pos)
		pos, chr, ori := make([]string, r), make([]string, r), make([]string, r)
		for j := 0; j < r; j++ {
			c := len(can[i].Pos[j])
			strp, strc, stro := make([]string, c), make([]string, c), make([]string, c)
			for k := 0; k < c; k++ {
				strp[k] = strconv.Itoa(can[i].Pos[j][k])
				strc[k] = strconv.Itoa(can[i].Chr[j][k] + 1)
				stro[k] = fmt.Sprintf("%.22f", can[i].Ori[j][k])
			}
			pos[j] = strings.Join(strp, ":")
			chr[j] = strings.Join(strc, ":")
			ori[j] = strings.Join(stro, ":")
		}
		chr[r-1] = "0" // ends with 0

		writer.Write(pos)
		writer.Write(chr)
		writer.Write(ori)
		writer.Flush()
	}
}
