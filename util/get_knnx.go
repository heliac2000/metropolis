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

import (
	"log"
	"math"
	"sync"
)

// Types
//
type annkdTree struct {
	pts                [][]float64
	dim, nPts, bktSize int
	pidx               []int
	root               interface{}
	bndBoxLo, bndBoxHi []float64
}

type annkdLeaf struct {
	nPts int
	bkt  []int
}

type annkdSplit struct {
	cutDim int
	cutVal float64
	cdBnds []float64
	child  []interface{}
}

//
// FNN/R/KNN.R: get.knnx() function
//
func GetKnnx(data, query [][]float64, k int) ([][]float64, [][]int) {

	if len(data[0]) != len(query[0]) {
		log.Fatalln("Number of columns must be same!")
	}
	if k > len(data) {
		log.Println("k should be less than sample size!")
	}

	return getKNNXkd(data, query, k)
}

//
// FNN/src/KNN_ANN.cpp: get_KNNX_kd() function
//
// Argments
//   K: Max. num of NN
//   d: Actual Dimension
//   n: Number of Base Data points
//   m: Number of Query Data points
//
func getKNNXkd(data, query [][]float64, K int) ([][]float64, [][]int) {
	tree := newAnnkdTree(data)

	var nn_dist [][]float64
	var nn_idx [][]int
	m := len(query)
	_, _ = Create2DimArray(&nn_dist, m, K), Create2DimArray(&nn_idx, m, K)

	var wg sync.WaitGroup
	wg.Add(m)
	for i := 0; i < m; i++ {
		go func(i int) {
			defer wg.Done()
			dist, index := annkSearch(tree, query[i], K)
			for j := 0; j < K; j++ {
				nn_dist[i][j] = math.Sqrt(dist[j])
				nn_idx[i][j] = index[j]
			}
		}(i)
	}
	wg.Wait()

	return nn_dist, nn_idx
}

//
// FNN/src/kd_tree.cpp: ANNkd_tree::ANNkd_tree()
//
func newAnnkdTree(data [][]float64) *annkdTree {
	if len(data) == 0 {
		return nil
	}

	n, dim := len(data), len(data[0])
	pidx := make([]int, n)
	for i := 0; i < n; i++ {
		pidx[i] = i
	}
	lo, hi := annEnclRect(data, pidx)
	loc, hic := CopyVector(lo).([]float64), CopyVector(hi).([]float64)
	root := rkdTree(data, pidx, n, loc, hic)

	return &annkdTree{
		pts: data, nPts: n, dim: dim, bktSize: 1,
		pidx: pidx, root: root, bndBoxLo: lo, bndBoxHi: hi,
	}
}

//
// FNN/src/kd_util.cpp: annEnclRect()
//
func annEnclRect(pa [][]float64, pidx []int) ([]float64, []float64) {
	n, dim := len(pa), len(pa[0])
	bndsLo, bndsHi := make([]float64, dim), make([]float64, dim)

	for d := 0; d < dim; d++ {
		loBnd, hiBnd := pa[pidx[0]][d], pa[pidx[0]][d]
		for i := 0; i < n; i++ {
			if v := pa[pidx[i]][d]; v < loBnd {
				loBnd = v
			} else if v > hiBnd {
				hiBnd = v
			}
		}
		bndsLo[d], bndsHi[d] = loBnd, hiBnd
	}

	return bndsLo, bndsHi
}

//
// FNN/src/kd_tree.cpp: ANNkd_ptr rkd_tree()
//
func rkdTree(pa [][]float64, pidx []int, n int, bndBoxLo, bndBoxHi []float64) interface{} {
	if n <= 1 {
		if n == 0 {
			return &annkdLeaf{nPts: 0, bkt: []int{0}}
		} else {
			return &annkdLeaf{nPts: n, bkt: pidx}
		}
	}

	cd, cv, nLo := slMidptSplit(pa, pidx, n, bndBoxLo, bndBoxHi)
	lv, hv := bndBoxLo[cd], bndBoxHi[cd]

	bndBoxHi[cd] = cv
	lo := rkdTree(pa, pidx, nLo, bndBoxLo, bndBoxHi)
	bndBoxHi[cd] = hv

	bndBoxLo[cd] = cv
	hi := rkdTree(pa, pidx[nLo:], n-nLo, bndBoxLo, bndBoxHi)
	bndBoxLo[cd] = lv

	// FNN/src/kd_tree.cpp: ANNkd_split::ANNkd_split()
	//
	return &annkdSplit{
		cutDim: cd,
		cutVal: cv,
		cdBnds: []float64{lv, hv},
		child:  []interface{}{lo, hi},
	}
}

//
// FNN/src/kd_split.cpp: void sl_midpt_split()
//
func slMidptSplit(pa [][]float64, pidx []int, n int, bndBoxLo, bndBoxHi []float64) (int, float64, int) {
	const ERR float64 = 0.001
	dim := len(pa[0])

	maxLength := bndBoxHi[0] - bndBoxLo[0]
	for d := 1; d < dim; d++ {
		length := bndBoxHi[d] - bndBoxLo[d]
		if length > maxLength {
			maxLength = length
		}
	}

	maxSpread := -1.0
	cutDim := 0
	for d := 0; d < dim; d++ {
		if (bndBoxHi[d] - bndBoxLo[d]) >= (1.0-ERR)*maxLength {
			if spr := annSpread(pa, pidx, n, d); spr > maxSpread {
				maxSpread = spr
				cutDim = d
			}
		}
	}

	iCutVal := (bndBoxLo[cutDim] + bndBoxHi[cutDim]) / 2.0
	min, max := annMinMax(pa, pidx, n, cutDim)
	cutVal := iCutVal
	if iCutVal < min {
		cutVal = min
	} else if iCutVal > max {
		cutVal = max
	}

	br1, br2 := annPlaneSplit(pa, pidx, n, cutDim, cutVal)
	var nLo int
	switch {
	case iCutVal < min:
		nLo = 1
	case iCutVal > max:
		nLo = n - 1
	case br1 > n/2:
		nLo = br1
	case br2 < n/2:
		nLo = br2
	default:
		nLo = n / 2
	}

	return cutDim, cutVal, nLo
}

//
// FNN/src/kd_util.cpp: annSpread()
//
func annSpread(pa [][]float64, pidx []int, n, d int) float64 {
	min, max := pa[pidx[0]][d], pa[pidx[0]][d]

	for i := 0; i < n; i++ {
		if c := pa[pidx[i]][d]; c < min {
			min = c
		} else if c > max {
			max = c
		}
	}

	return (max - min)
}

//
// FNN/src/kd_util.cpp: annMinMax()
//
func annMinMax(pa [][]float64, pidx []int, n, d int) (float64, float64) {
	min, max := pa[pidx[0]][d], pa[pidx[0]][d]

	for i := 0; i < n; i++ {
		if c := pa[pidx[i]][d]; c < min {
			min = c
		} else if c > max {
			max = c
		}
	}

	return min, max
}

//
// FNN/src/kd_util.cpp: annPlaneSplit()
//
func annPlaneSplit(pa [][]float64, pidx []int, n, d int, cv float64) (int, int) {
	l, r := 0, n-1

	for {
		for l < n && pa[pidx[l]][d] < cv {
			l++
		}
		for r >= 0 && pa[pidx[r]][d] >= cv {
			r--
		}
		if l > r {
			break
		}
		pidx[l], pidx[r] = pidx[r], pidx[l]
		l, r = l+1, r-1
	}

	br1, r := l, n-1
	for {
		for l < n && pa[pidx[l]][d] <= cv {
			l++
		}
		for r >= br1 && pa[pidx[r]][d] > cv {
			r--
		}
		if l > r {
			break
		}
		pidx[l], pidx[r] = pidx[r], pidx[l]
		l, r = l+1, r-1
	}

	return br1, l
}

//
// Types
//
type mkNode struct {
	key  float64
	info int
}

type annMinK struct {
	k, n int
	mk   []mkNode
}

//
// FNN/src/kd_search.cpp: shared variables
//
type annSharedVars struct {
	annKdDim     int
	annKdQ       []float64
	annKdPts     [][]float64
	annKdPointMK *annMinK
}

//
// FNN/src/kd_search.cpp: ANNkd_tree::annkSearch()
//
func annkSearch(tree *annkdTree, q []float64, k int) ([]float64, []int) {
	if k > tree.nPts {
		log.Fatalln("Requesting more near neighbors than data points")
	}

	sharedVars := &annSharedVars{
		annKdDim:     tree.dim,
		annKdQ:       q,
		annKdPts:     tree.pts,
		annKdPointMK: newAnnMinK(k),
	}

	dist := annBoxDistance(q, tree.bndBoxLo, tree.bndBoxHi, tree.dim)
	switch tree.root.(type) {
	case *annkdSplit:
		tree.root.(*annkdSplit).AnnSearch(dist, sharedVars)
	case *annkdLeaf:
		tree.root.(*annkdLeaf).AnnSearch(dist, sharedVars)
	}

	dd, nn_idx := make([]float64, k), make([]int, k)
	for i := 0; i < k; i++ {
		dd[i] = sharedVars.annKdPointMK.ithSmallestKey(i)
		nn_idx[i] = sharedVars.annKdPointMK.ithSmallestInfo(i)
	}

	return dd, nn_idx
}

func (node *annkdSplit) AnnSearch(boxDist float64, sharedVars *annSharedVars) {
	cutDiff := sharedVars.annKdQ[node.cutDim] - node.cutVal
	if cutDiff < 0 {
		switch node.child[0].(type) {
		case *annkdSplit:
			node.child[0].(*annkdSplit).AnnSearch(boxDist, sharedVars)
		case *annkdLeaf:
			node.child[0].(*annkdLeaf).AnnSearch(boxDist, sharedVars)
		}

		boxDiff := node.cdBnds[0] - sharedVars.annKdQ[node.cutDim]
		if boxDiff < 0 {
			boxDiff = 0
		}
		boxDist += cutDiff*cutDiff - boxDiff*boxDiff

		if boxDist < sharedVars.annKdPointMK.maxKey() {
			switch node.child[1].(type) {
			case *annkdSplit:
				node.child[1].(*annkdSplit).AnnSearch(boxDist, sharedVars)
			case *annkdLeaf:
				node.child[1].(*annkdLeaf).AnnSearch(boxDist, sharedVars)
			}
		}
	} else {
		switch node.child[1].(type) {
		case *annkdSplit:
			node.child[1].(*annkdSplit).AnnSearch(boxDist, sharedVars)
		case *annkdLeaf:
			node.child[1].(*annkdLeaf).AnnSearch(boxDist, sharedVars)
		}

		boxDiff := sharedVars.annKdQ[node.cutDim] - node.cdBnds[1]
		if boxDiff < 0 {
			boxDiff = 0
		}
		boxDist += cutDiff*cutDiff - boxDiff*boxDiff

		if boxDist < sharedVars.annKdPointMK.maxKey() {
			switch node.child[0].(type) {
			case *annkdSplit:
				node.child[0].(*annkdSplit).AnnSearch(boxDist, sharedVars)
			case *annkdLeaf:
				node.child[0].(*annkdLeaf).AnnSearch(boxDist, sharedVars)
			}
		}
	}

	return
}

func (node *annkdLeaf) AnnSearch(boxDist float64, sharedVars *annSharedVars) {
	minDist, dist := sharedVars.annKdPointMK.maxKey(), 0.0

	for i := 0; i < node.nPts; i++ {
		pp, d := sharedVars.annKdPts[node.bkt[i]], 0
		for ; d < sharedVars.annKdDim; d++ {
			dist += (sharedVars.annKdQ[d] - pp[d]) * (sharedVars.annKdQ[d] - pp[d])
			if dist > minDist {
				break
			}
		}
		if d >= sharedVars.annKdDim {
			sharedVars.annKdPointMK.insert(dist, node.bkt[i])
			minDist = sharedVars.annKdPointMK.maxKey()
		}
	}

	return
}

//
// FNN/src/kd_util.cpp: ANNdist annBoxDistance()
//
func annBoxDistance(q, lo, hi []float64, dim int) float64 {
	dist := 0.0
	for d, t := 0, 0.0; d < dim; d++ {
		if q[d] < lo[d] {
			t = lo[d] - q[d]
		} else if q[d] > hi[d] {
			t = q[d] - hi[d]
		}
		dist += t * t
	}

	return dist
}

//
// FNN/src/pr_queue_k.h: ANNmin_k::ANNmin_k()
//
func newAnnMinK(max int) *annMinK {
	return &annMinK{
		n: 0, k: max, mk: make([]mkNode, max+1),
	}
}

func (mink *annMinK) maxKey() float64 {
	if mink.n == mink.k {
		return mink.mk[mink.k-1].key
	}
	return math.MaxFloat64
}

func (mink *annMinK) ithSmallestKey(i int) float64 {
	if i < mink.n {
		return mink.mk[i].key
	}
	return math.MaxFloat64
}

func (mink *annMinK) ithSmallestInfo(i int) int {
	if i < mink.n {
		return mink.mk[i].info
	}
	return -1
}

func (mink *annMinK) insert(kv float64, inf int) {
	i := mink.n
	for ; i > 0; i-- {
		if mink.mk[i-1].key > kv {
			mink.mk[i] = mink.mk[i-1]
		} else {
			break
		}
	}

	mink.mk[i].key, mink.mk[i].info = kv, inf
	if mink.n < mink.k {
		mink.n++
	}

	return
}
