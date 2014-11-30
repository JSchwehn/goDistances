package goDistances

import (
	"errors"
	"fmt"
	"math"
)

// LpNormDistance takes one float64 slices and one float64 value.
// It will return a float for the distance between the two values.
type LpNormDistance struct{}

func (e2d LpNormDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 2 {
		return -1., errors.New(fmt.Sprintf("Wrong parameter count. Needed 2 got %d", len(params)))
	}
	vector := params[0].([]float64)
	p := params[1].(float64)

	d := 0.0
	for _, i := range vector {
		d += math.Pow(math.Abs(i), p)
	}
	return math.Pow(d, 1/p), nil
}
