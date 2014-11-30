package goDistances

import(
	"fmt"
	"math"
	"errors"
)

type minkowskiDistance struct {}
// p norm distance
func (md minkowskiDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 3 {
		return -1, errors.New(fmt.Sprintf("Wrong parameter count. Needed 3 got %d", len(params)))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)
	p := params[2].(float64)

	if len(v1) != len(v2) {
		return -1., errors.New("Type mismatch - size of vector 1 is not equal to vector 2")
	}
	d := 0.0
	for i := range v1 {
		d += math.Pow(math.Abs(v1[i]-v2[i]),p)
	}
	return math.Pow(d,1/p), nil
}
