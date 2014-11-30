package goDistances

import(
	"fmt"
	"math"
	"errors"
)

// MinkowskiDistanceWeighted takes three float64 slices which have to have the same size.
// The first two slices are the vectors and the third slice is the weight table. The forth
// parameter is a float64 value.
// It will return a float for the distance between the two slices.
type MinkowskiDistanceWeighted struct {}
// p norm distance
func (mwd MinkowskiDistanceWeighted) Distance(params ...interface{}) (float64, error) {
	if len(params) != 4 {
		return -1, errors.New(fmt.Sprintf("Wrong parameter count. Needed 4 got %d", len(params)))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)
	weights := params[2].([]float64)
	p := params[3].(float64)

	if len(v1) != len(v2) || len(v1) != len(weights) {
		return -1., errors.New("Type mismatch - size of vector 1 is not equal to vector 2 or not equal to the wights")
	}
	d := 0.0
	for i := range v1 {
		d += weights[i] * math.Pow(math.Abs(v1[i]-v2[i]),p)
	}
	return math.Pow(d,1/p), nil
}
