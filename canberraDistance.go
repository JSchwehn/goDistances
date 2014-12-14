package goDistances

import (
	"fmt"
	"math"
)

// CanberraDistance calculates a canberra distance.
// see http://en.wikipedia.org/wiki/Canberra_distance
// see http://reference.wolfram.com/language/ref/CanberraDistance.html
type CanberraDistance struct{}

// Distance takes two float64 slices which have to have the same size. It will return a float for the distance between the two slices.
func (cd CanberraDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 2 {
		return -1, fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)

	d := 0.
	for i := range v1 {
		d += (math.Abs(v1[i]-v2[i]) / (math.Abs(v1[i]) + math.Abs(v2[i])))
	}
	return d, nil
}
