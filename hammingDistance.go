package goDistances

import(
	"fmt"
	"errors"
)

type hammingDistance struct {}

func(hd hammingDistance) Distance(params ...interface {}) (float64,error){
	if len(params) != 2 {
		return -1, errors.New(fmt.Sprintf("Wrong parameter count. Needed 3 got %d", len(params)))
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
