package goDistances

import(
	"fmt"
	"math"
	"errors"
)

// see http://people.revoledu.com/kardi/tutorial/Similarity/BrayCurtisDistance.html
// http://www.wolframalpha.com/input/?i=Bray+Curtis+Distance
type brayCurtisDistance struct {}

func(bcd brayCurtisDistance) Distance(params ...interface {}) (float64,error){
	if len(params) != 2 {
		return -1, errors.New(fmt.Sprintf("Wrong parameter count. Needed 3 got %d", len(params)))
	}
	v1 := params[0].([]float64)
	v2 := params[1].([]float64)

	numerator, denominator := 0.0, 0.0

	for i := range v1 {
		numerator += math.Abs(v1[i]-v2[i])
		denominator += math.Abs(v1[i]+v2[i])
	}

	return numerator / denominator, nil
}
