package utils

import (
	"errors"
	"math"
)

func GetEuclideanDistance(firstPoint []float64, secondPoint []float64) (float64, error) {
	if len(firstPoint) != len(secondPoint) {
		return 0, errors.New("points must have the same dimensions")
	}
	sum := 0.0
	for i := 0; i < len(firstPoint); i++ {
		sum += math.Pow(firstPoint[i]-secondPoint[i], 2)
	}
	return math.Sqrt(sum), nil
}
