package goDistances

import(
	"fmt"
	"math"
	"errors"
)

//canberraDistance takes two float64 slices which have to have the same size. It will return a float for the distance between the two slices.
// see http://en.wikipedia.org/wiki/Canberra_distance
// see http://reference.wolfram.com/language/ref/CanberraDistance.html
type CanberraDistance struct {}

func(cd CanberraDistance) Distance(params ...interface {}) (float64,error){
	if len(params) != 2 {
		return -1, errors.New(fmt.Sprintf("Wrong parameter count. Needed 3 got %d", len(params)))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)

	d := 0.
	for i := range v1 {
		d += (math.Abs(v1[i]-v2[i]) / (math.Abs(v1[i])+math.Abs(v2[i])))
	}
	return d, nil
}
