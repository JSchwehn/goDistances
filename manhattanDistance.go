package goDistances

import (
	"fmt"
	"errors"
	"math"
)

// EuclideanDistance calculates the manhattan distance
type ManhattanDistance struct {}

// Distance takes two float64 slices which have to have the same size.
// It will return a float for the distance between the two slices.
func (mD ManhattanDistance) Distance(params ...interface{}) (float64,error) {
	if len(params) != 2 {
		return -1., errors.New(fmt.Sprintf("Wrong parameter count. Needed 2 got %d", len(params)))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)

	if len(v1) != len(v2) {
		return -1, errors.New("Type mismatch - size of vector 1 is not equal to vector 2")
	}
	d := 0.0
	for i := range v1 {
		d += math.Abs(v1[i]-v2[i])
	}
	return d, nil
}
