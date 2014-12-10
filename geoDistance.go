package goDistances

import (
	"fmt"
//	"math"
)

// GeoDistance calculates the distance
type GeoDistance struct {}
//type GeoCoordinate struct {
//	degree  float64
//	minutes float64
//	seconds float64
//	ns string // n(orth) or (s)outh
//	ew string // e(ast) or w(est)
//}

// Distance takes two float64 slices which have to have two elements each.
// The first parameter of the slice  will be considered as north angle and the second
// one will be considered as east angle.
// It will return a float for the distance between the two slices.
func (mG GeoDistance) Distance(params ...interface{}) (float64,error) {
	if len(params) != 3 {
		return -1.,  fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	d1 := params[0].([]float64) // latitude *longitute
	d2 := params[1].([]float64) //
	c := params[2].(float64)

	fmt.Printf("We got %v %v %v  \n",d1, d2, c)
	d := 0.0

	return d, nil
}

func (mG GeoDistance) ConvertDMSToDegrees(dmsLati, dmsLong string) ([]float64){
	fmt.Printf("DMS: %v %v ",dmsLati, dmsLong)
	return []float64{1.,2.}
}
