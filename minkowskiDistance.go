package goDistances

import(
	"fmt"
	"math"
)
// MinkowskiDistance calculates the Minkowski distance
type MinkowskiDistance struct {}

// Distance takes two float64 slices which have to have the same size and a float64 value.
// It will return a float for the distance between the two slices.
func (md MinkowskiDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 3 {
		return -1,  fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)
	p := params[2].(float64)

	if len(v1) != len(v2) {
		return -1.,  fmt.Errorf("type mismatch - size of vector 1 is not equal to vector 2")
	}
	d := 0.0
	for i := range v1 {
		d += math.Pow(math.Abs(v1[i]-v2[i]),p)
	}
	return math.Pow(d,1/p), nil
}
