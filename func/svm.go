//
// svm.go
//

package functions

import (
	"fmt"
	"log"
	"math"
)

import . "../util"

// R's SVM object
//
type XScaleObjct struct {
	ScaledCenter []float64 `json:"scaled:center"`
	ScaledScale  []float64 `json:"scaled:scale"`
}

type Svm struct {
	Type           int
	Kernel         int
	Cost           float64
	Degree         int
	Gamma          float64
	Coef0          float64
	Nu             float64
	Epsilon        float64
	Sparse         bool
	Scaled         []bool
	XScale         XScaleObjct `json:"x.scale"`
	NClasses       int
	Levels         []string
	TotNsv         int `json:"tot.nsv"`
	Nsv            []int
	Labels         []int
	Sv             [][]float64
	Index          []int
	Rho            []float64
	CompProb       bool
	Coefs          [][]float64
	Fitted         []string
	DecisionValues [][]float64 `json:"decision.values"`
}

// R's predict generic function
//
func (svm *Svm) Predict(newd [][]float64) string {

	if svm.TotNsv < 1 {
		log.Fatalln("Model is empty!")
	}

	rowNs := make([]int, len(newd))
	for i := 0; i < len(newd); i++ {
		rowNs[i] = i
	}

	if Any(svm.Scaled) {
		r, c := len(newd), len(newd[0])
		var newData [][]float64
		Create2DimArray(&newData, r, c)
		n := 0
		for j := 0; j < c; j++ {
			if svm.Scaled[j] {
				for i := 0; i < r; i++ {
					newData[i][n] = newd[i][j]
				}
				n++
			}
		}
		if n < c {
			for i := 0; i < r; i++ {
				newData[i] = newData[i][:n]
			}
		}

		newd = Scale(newData, svm.XScale.ScaledCenter, svm.XScale.ScaledScale)
	}

	if len(svm.Sv[0]) != len(newd[0]) {
		log.Fatalln("Test data does not match model !")
	}

	// R is 1-base, Golang is 0-base
	return svm.Levels[svm.SvmPredict(newd)[0]-1]
}

// [R] e1071/src/svm.h
//
type svmNode struct {
	index int
	value float64
}

type svmParameter struct {
	svmType    int
	kernelType int
	degree     int     // for poly
	gamma      float64 // for poly/rbf/sigmoid
	coef0      float64 // for poly/sigmoid

	// these are for training only
	cacheSize   float64   // in MB
	eps         float64   // stopping criteria
	c           float64   // for C_SVC, EPSILON_SVR and NU_SVR
	nrWeight    int       // for C_SVC
	weightLabel []int     // for C_SVC
	weight      []float64 // for C_SVC
	nu          float64   // for NU_SVC, ONE_CLASS, and NU_SVR
	p           float64   // for EPSILON_SVR
	shrinking   int       // use the shrinking heuristics
	probability bool      // do probability estimates

}

type svmModel struct {
	svmParam  svmParameter // parameter
	nrClass   int          // number of classes, = 2 in regression/one class svm
	l         int          // total #SV
	svNodes   [][]svmNode  //struct svm_node **SV // SVs (SV[l])
	svCoef    [][]float64  // coefficients for SVs in decision functions (sv_coef[k-1][l])
	rho       []float64    // constants in decision functions (rho[k*(k-1)/2])
	probA     float64      // pariwise probability information
	probB     float64
	svIndices []int // sv_indices[0,...,nSV-1] are values in [1,...,num_traning_data] to indicate SVs in the training set

	// for classification only
	label []int // label of each class (label[k])

	// number of SVs for each class (nSV[k])
	// nSV[0] + nSV[1] + ... + nSV[k-1] = l
	nSV []int

	// 1 if svm_model is created by svm_load_model
	// 0 if svm_model is created by svm_train
	freeSv int
}

// [R] e1071/src/Rsvm.c: svmpredict() function
//
func (svm *Svm) SvmPredict(newd [][]float64) []int {
	// set up model
	m := &svmModel{
		l:       svm.TotNsv,
		nrClass: svm.NClasses,
		svCoef: func() [][]float64 {
			return Transpose(svm.Coefs).([][]float64)
		}(),
		svNodes: sparsify(svm.Sv), // [R] svm_model$sparse == FALSE
		rho:     svm.Rho,
		probA:   0, // [R] svm_model$compprob == FALSE
		probB:   0, // [R] svm_model$compprob == FALSE
		label:   svm.Labels,
		nSV:     svm.Nsv,
		freeSv:  1,
		svmParam: svmParameter{
			svmType:     svm.Type,
			kernelType:  svm.Kernel,
			degree:      svm.Degree,
			gamma:       svm.Gamma,
			coef0:       svm.Coef0,
			probability: svm.CompProb,
		},
	}

	// create sparse training matrix
	train := sparsify(newd) // [R] sparse == FALSE

	// call svm-predict-function for each x-row, possibly using
	// probability estimator, if requested
	//
	// [R] probability == FALSE
	//
	ret := make([]int, len(newd))
	for i := 0; i < len(newd); i++ {
		ret[i] = m.svmPredictValues(train[i])
	}

	fmt.Println(ret)

	return ret
}

// [R] e1071/src/Rsvm.c: sparsify() function
//
func sparsify(x [][]float64) [][]svmNode {
	r, c := len(x), len(x[0])
	sparse := make([][]svmNode, r)
	for i := 0; i < r; i++ {
		// allocate memory for column elements
		sparse[i] = make([]svmNode, 0, c)

		// set column elements
		for ii := 0; ii < c; ii++ {
			// determine nr. of non-zero elements
			if x[i][ii] != 0 {
				sparse[i] = append(sparse[i],
					svmNode{index: ii + 1, value: x[i][ii]})
			}
		}
	}

	return sparse
}

// [R] e1071/src/svm.cpp: svm_predict() function
//
func (model *svmModel) svmPredictValues(x []svmNode) int {
	var decValues []float64
	decValues = make([]float64, model.nrClass*(model.nrClass-1)/2)

	kValue := make([]float64, model.l)
	for i := 0; i < model.l; i++ {
		kValue[i] = kFunction(x, model.svNodes[i], &model.svmParam)
	}

	start := make([]int, model.nrClass)
	for i := 1; i < model.nrClass; i++ {
		start[i] = start[i-1] + model.nSV[i-1]
	}
	vote := make([]int, model.nrClass)

	p := 0
	for i := 0; i < model.nrClass; i++ {
		for j := i + 1; j < model.nrClass; j++ {
			sum := 0.0
			si, sj := start[i], start[j]
			ci, cj := model.nSV[i], model.nSV[j]
			for k := 0; k < ci; k++ {
				sum += model.svCoef[j-1][si+k] * kValue[si+k]
			}
			for k := 0; k < cj; k++ {
				sum += model.svCoef[i][sj+k] * kValue[sj+k]
			}
			sum -= model.rho[p]
			decValues[p] = sum
			if decValues[p] > 0 {
				vote[i]++
			} else {
				vote[j]++
			}
			p++
		}
	}

	voteMaxIdx := 0
	for i := 1; i < model.nrClass; i++ {
		if vote[i] > vote[voteMaxIdx] {
			voteMaxIdx = i
		}
	}

	return model.label[voteMaxIdx]
}

// kernel_type
const (
	LINEAR int = iota
	POLY
	RBF
	SIGMOID
	PRECOMPUTED
)

// [R] e1071/src/svm.cpp: Kernel::k_function() method(static)
//
func kFunction(x []svmNode, y []svmNode, param *svmParameter) float64 {
	switch param.kernelType {
	case LINEAR:
		return dot(x, y)
	case POLY:
		return powi(param.gamma*dot(x, y)+param.coef0, param.degree)
	case RBF:
		return kernelRBF(x, y, param.gamma)
	case SIGMOID:
		return math.Tanh(param.gamma*dot(x, y) + param.coef0)
	case PRECOMPUTED: //x: test (validation), y: SV
		return x[int(y[0].value)].value
	default:
		return 0 // Unreachable
	}
}

func kernelRBF(x []svmNode, y []svmNode, gamma float64) float64 {
	sum, i, j := 0.0, 0, 0
	for i < len(x) && j < len(y) {
		if x[i].index == y[j].index {
			sum += (x[i].value - y[j].value) * (x[i].value - y[j].value)
			i, j = i+1, j+1
		} else {
			if x[i].index > y[j].index {
				sum += y[j].value * y[j].value
				j++
			} else {
				sum += x[i].value * x[i].value
				i++
			}
		}
	}

	for ; i < len(x); i++ {
		sum += x[i].value * x[i].value
	}
	for ; j < len(y); j++ {
		sum += y[j].value * y[j].value
	}

	return math.Exp(-gamma * sum)
}

// [R] e1071/src/svm.cpp: Kernel::dot() method(static)
//
func dot(px []svmNode, py []svmNode) float64 {
	sum := 0.0
	for i, j := 0, 0; i < len(px) && j < len(py); {
		if px[i].index == py[j].index {
			sum += px[i].value * py[j].value
			i, j = i+1, j+1
		} else {
			if px[i].index > py[j].index {
				j++
			} else {
				i++
			}
		}
	}

	return sum
}

// [R] e1071/src/svm.cpp: inline powi()
//
func powi(base float64, times int) float64 {
	tmp, ret := base, 1.0

	for t := times; t > 0; t /= 2 {
		if t%2 == 1 {
			ret *= tmp
		}
		tmp = tmp * tmp
	}

	return ret
}

// svm_type
// const (
// 	C_SVC int = iota
// 	NU_SVC
// 	ONE_CLASS
// 	EPSILON_SVR
// 	NU_SVR
// )
//
// func (model *svmModel) svmPredictValues(x []svmNode) float64 {
// 	var decValues []float64
//
// 	if model.svmParam.svmType == ONE_CLASS ||
// 		model.svmParam.svmType == EPSILON_SVR ||
// 		model.svmParam.svmType == NU_SVR {
// 		decValues = make([]float64, 1)
// 	} else {
// 		decValues = make([]float64, model.nrClass*(model.nrClass-1)/2)
// 	}
//
// 	_ = decValues
// 	return 0
// }
