package goDistances

import (
	"math"
	"testing"
)
var v1 = []float64{4.0, 1.0, -2.0}
var v2 = []float64{2.0, 3.0, -1.0}
// Euclid distance success
func TestEclid(t *testing.T) {
	const out1 = 3.0
	const precision = 0.00001
	euclid := new(EuclideanDistance)
	distance, err := euclid.Distance(v1, v2)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !isEqual(distance, out1, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v with an precision of %v", distance, out1, precision)
	}
}

// Euclid distance type mismatch
func TestEclidTypeMismatch(t *testing.T) {
	v1 := []float64{4, 1}
	euclid := new(EuclideanDistance)
	_, err := euclid.Distance(v1, v2)
	if err == nil {
		t.Errorf("Expected an error, but did't received one")
	}
}

// Euclid distance success
func TestEclidSquare(t *testing.T) {
	const out1 = 9.0
	const precision = 0.00001
	euclid2 := new(EuclideanDistanceSquare)
	distance, err := euclid2.Distance(v1, v2)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !isEqual(distance, out1, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out1,precision)
	}
}

//LPNorm Distance
func TestNorm(t *testing.T) {
	const out = 4.05885
	const precision = 0.00001
	p 	:= 4.1
	norm := new(LpNormDistance)
	distance, err := norm.Distance(v1, p)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}

func TestNormTypeMismatch(t *testing.T) {
	v2 := []float64{2, 3, -1}
	const p = 4.1
	norm := new(LpNormDistance)
	_, err := norm.Distance(v1, v2, p)
	if err == nil {
		t.Errorf("Expected an error, but did't received one")
	}
}

func TestManhattanDistance(t *testing.T) {
	const out = 5.0
	manhattan := new(ManhattanDistance)
	distance, err := manhattan.Distance(v1,v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if distance != out {
		t.Errorf("Distance is wrong. Wanted %v - got %v",out, distance)
	}
}

func TestMinkowskiDistanceEclideanLike(t *testing.T) {
	const out = 3.0
	minkowskiDistance := new(MinkowskiDistance)
	distance, err := minkowskiDistance.Distance(v1,v2,2.0)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if distance != out {
		t.Errorf("Distance is wrong. Wanted %v - got %v",out, distance)
	}
}
func TestMinkowskiDistanceManhattanLike(t *testing.T) {
	const out = 5.0
	minkowskiDistance := new(MinkowskiDistance)
	distance, err := minkowskiDistance.Distance(v1,v2,1.0)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if distance != out {
		t.Errorf("Distance is wrong. Wanted %v - got %v",out, distance)
	}
}
func TestMinkowskiDistanceWeighted(t *testing.T) {
	const out = 4.24264
	const precision = 0.00001
	weights := []float64{2.,2.,2.}
	minkowskiDistanceWeighted := new(MinkowskiDistanceWeighted)
	distance, err := minkowskiDistanceWeighted.Distance(v1,v2,weights,2.0)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestChebyshevDistance(t *testing.T) {
	const out = 2.
	const precision = 0.00001
	chebyshevDistance := new(ChebyshevDistance)
	distance, err := chebyshevDistance.Distance(v1,v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestHammingDistance(t *testing.T) {
	const out = 3.
	const precision = 0.00001
	hammingDistance := new(HammingDistance)
	distance, err := hammingDistance.Distance(v1,v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestBrayCurtisDistance(t *testing.T) {
	const out       = 0.384615
	const precision = 0.000001
	brayCurtis := new(BrayCurtisDistance)
	distance, err := brayCurtis.Distance(v1,v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestCanberraDistance(t *testing.T) {
	const out       = 1.166667
	const precision = 0.000001
	canberraDistance := new(CanberraDistance)
	distance, err := canberraDistance.Distance(v1,v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v",err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}

func isEqual(f1, f2 float64, tolerance float64) bool {
	if math.Abs((f1-f2)) <= tolerance {
		return true
	}
	return false
}
