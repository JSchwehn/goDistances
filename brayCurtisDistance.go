package goDistances

import (
	"fmt"
	"math"
)

// BrayCurtisDistance calculates a Bray Curtis distance.
// see http://people.revoledu.com/kardi/tutorial/Similarity/BrayCurtisDistance.html
// http://www.wolframalpha.com/input/?i=Bray+Curtis+Distance
type BrayCurtisDistance struct{}

// Distance takes two float slices as vectors and returns a float as the distance between the two vectors.
func (bcd BrayCurtisDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 2 {
		return -1, fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)

	numerator, denominator := 0.0, 0.0

	for i := range v1 {
		numerator += math.Abs(v1[i] - v2[i])
		denominator += math.Abs(v1[i] + v2[i])
	}

	return numerator / denominator, nil
}
