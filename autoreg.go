package goautoreg

import (
	"errors"
)

type ARFit func(data []float64, order int) ([]float64, error)

func normalizedCorrelation(signal []float64, maxLag int) ([]float64, error) {
	result, err := correlation(signal, maxLag)
	if err != nil {
		return nil, err
	}
	return vecScalarDiv(result, result[0]), nil
}

func correlation(signal []float64, maxLag int) ([]float64, error) {
	signalLength := len(signal)
	if maxLag >= signalLength {
		return nil, errors.New("maxLag cannot be greater, or equal than the signal length")
	}

	results := make([]float64, maxLag+1)
	for lag := 0; lag <= maxLag; lag++ {
		results[lag] = vecMul(signal[:signalLength-lag], signal[lag:])
	}
	return results, nil
}

func vecMul(v1 []float64, v2 []float64) float64 {
	sum := 0.0
	for i := range v1 {
		sum += v1[i] * v2[i]
	}
	return sum
}

func vecScalarDiv(s []float64, d float64) []float64 {
	result := make([]float64, len(s))
	for i, v := range s {
		result[i] = v / d
	}
	return result
}
