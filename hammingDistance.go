package goDistances

import(
	"fmt"
)


// HammingDistance calculates the hamming distance
type HammingDistance struct {}

// Distance takes two float64 slices which have to have the same size.
// It will return a float for the distance between the two slices.
func(hd HammingDistance) Distance(params ...interface {}) (float64,error){
	if len(params) != 2 {
		return -1,  fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)

	d := 0
	for i := range v1 {
		if v1[i] != v2[i] {
			d++
		}
	}
	return float64(d), nil
}
