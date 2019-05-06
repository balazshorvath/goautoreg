package goautoreg

import (
	"github.com/gonum/matrix/mat64"
	"github.com/pkg/errors"
)

// NewYuleWalkerARFit returns a function that is capable of calculating the coefficients of the autoregressive model
// with the Yule-Walker equations .
//
func NewYuleWalkerARFit() ARFit {
	return yuleWalkerARFit
}

func buildToeplitz(data []float64) *mat64.Dense {
	dataLen := len(data)
	reversed := make([]float64, dataLen)
	for i, v := range data {
		reversed[(dataLen-1)-i] = v
	}
	matrix := mat64.NewDense(dataLen, dataLen, nil)
	for i := 0; i < dataLen; i++ {
		row := append([]float64(nil), reversed[dataLen-1-i:dataLen-1]...)
		matrix.SetRow(i, append(row, data[:dataLen-i]...))
	}
	return matrix
}

// Fit Uses the Yule-Walker equations to calculate the coefficients of the autoregressive model.
//
func yuleWalkerARFit(data []float64, order int) ([]float64, error) {
	corr, err := normalizedCorrelation(data, order)
	if err != nil {
		return nil, errors.Wrap(err, "fit failed")
	}
	// Do r = R * Phi
	toeplitz := buildToeplitz(corr[:len(corr)-1])
	corrVector := mat64.NewVector(order, corr[1:])

	err = toeplitz.Inverse(toeplitz)
	if err != nil {
		return nil, errors.Wrap(err, "fit failed")
	}
	corrVector.MulVec(toeplitz, corrVector)

	return corrVector.RawVector().Data, nil
}
