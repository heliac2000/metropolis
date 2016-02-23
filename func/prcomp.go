//
// prcomp.go
//

package functions

import . "../util"

// R's prcomp object
//
type PrComp struct {
	Sdev     []float64
	Rotation [][]float64
	Center   []float64
	Scale    bool
	X        [][]float64
}

// R's predict generic function
//
func (prc *PrComp) Predict(newData [][]float64) [][]float64 {
	// if len(newData[0]) != len(prc.Rotation) {
	// 	log.Fatalln("'newdata' does not have the correct number of columns.")
	// }

	var sc [][]float64
	if !prc.Scale {
		sc = ScaleWithoutScaling(newData, prc.Center)
	}

	return MatrixMultiplyFloat(sc[0:1], prc.Rotation)
}

/*
[R]
> getAnywhere("predict.prcomp")
A single object matching ‘predict.prcomp’ was found
It was found in the following places
  registered S3 method for predict from namespace stats
  namespace:stats
with value

function (object, newdata, ...)
{
    if (missing(newdata)) {
        if (!is.null(object$x))
            return(object$x)
        else stop("no scores are available: refit with 'retx=TRUE'")
    }
    if (length(dim(newdata)) != 2L)
        stop("'newdata' must be a matrix or data frame")
    nm <- rownames(object$rotation)
    if (!is.null(nm)) {
        if (!all(nm %in% colnames(newdata)))
            stop("'newdata' does not have named columns matching one or more of the original columns")
        newdata <- newdata[, nm, drop = FALSE]
    }
    else {
        if (NCOL(newdata) != NROW(object$rotation))
            stop("'newdata' does not have the correct number of columns")
    }
    scale(newdata, object$center, object$scale) %*% object$rotation
}
*/
