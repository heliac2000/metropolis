package util

import "math"

// R's colMeans function
//
func ColMeans(m [][]float64) []float64 {
	r, c := len(m), len(m[0])
	means := make([]float64, c)

	for i := 0; i < c; i++ {
		sum := 0.0
		for j := 0; j < r; j++ {
			sum += m[j][i]
		}
		means[i] = sum / float64(r)
	}

	return means
}

// R's apply(m, 2, sd)
//
// From 'help(sd)', "Like 'var' this uses denominator n - 1."
//
func ColSd(m [][]float64) []float64 {
	r, c := len(m), len(m[0])
	sd := make([]float64, c)

	means := ColMeans(m)
	for i := 0; i < c; i++ {
		variance := 0.0
		for j := 0; j < r; j++ {
			variance += (m[j][i] - means[i]) * (m[j][i] - means[i])
		}
		sd[i] = math.Sqrt(variance / float64(r-1))
	}

	return sd
}

// R's scale function
//
func Scale(m [][]float64, center []float64, scale []float64) [][]float64 {
	var scaled [][]float64
	Copy2DimArray(&scaled, m)

	r, c := len(m), len(m[0])
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			scaled[j][i] = (scaled[j][i] - center[i]) / scale[i]
		}
	}

	return scaled
}
