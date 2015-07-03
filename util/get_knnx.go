//
// get_knnx.go
//

//
// FNN package/CRAN
// http://cran.r-project.org/web/packages/FNN/index.html
//
// Fast Nearest Neighbor Search Algorithms and Applications.
//

package util

import "log"

func GetKnnx(data, query [][]float64, k int) ([][]float64, [][]float64) {
	n, m := len(data), len(query)
	d, p := len(data[0]), len(query[0])

	if d != p {
		log.Fatalln("Number of columns must be same!")
	}
	if k > n {
		log.Println("k should be less than sample size!")
	}

	return get_KNNX_kd(Transpose(data), Transpose(query), k, d, n, m)
}

func get_KNNX_kd(data, query [][]float64, k, d, n, m int) ([][]float64, [][]float64) {
	return [][]float64{{}}, [][]float64{{}}
}
