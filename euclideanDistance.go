package goDistances

import (
	"fmt"
	"math"
)

// EuclideanDistance calculates an euclidean distance
// http://en.wikipedia.org/wiki/Euclidean_distance
type EuclideanDistance struct{}

// Distance takes two float64 slices which have to have the same size.
// It will return a float for the distance between the two slices.
func (e2d EuclideanDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 2 {
		return 0., fmt.Errorf("wrong parameter count. Needed 2 got %d", len(params))
	}

	vector0 := params[0].([]float64)
	vector1 := params[1].([]float64)

	if len(vector0) != len(vector1) {
		return -1, fmt.Errorf("type mismatch - size of vector1 is not equal to vector2")
	}
	d := 0.0
	for i := range vector0 {
		d += (vector0[i] - vector1[i]) * (vector0[i] - vector1[i])
	}
	return math.Sqrt(d), nil
}
