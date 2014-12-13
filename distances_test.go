package goDistances

import (
	"math"
	"testing"
	//	"fmt"
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
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out1, precision)
	}
}

//LPNorm Distance
func TestNorm(t *testing.T) {
	const out = 4.05885
	const precision = 0.00001
	p := 4.1
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
	distance, err := manhattan.Distance(v1, v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if distance != out {
		t.Errorf("Distance is wrong. Wanted %v - got %v", out, distance)
	}
}

func TestMinkowskiDistanceEclideanLike(t *testing.T) {
	const out = 3.0
	minkowskiDistance := new(MinkowskiDistance)
	distance, err := minkowskiDistance.Distance(v1, v2, 2.0)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if distance != out {
		t.Errorf("Distance is wrong. Wanted %v - got %v", out, distance)
	}
}
func TestMinkowskiDistanceManhattanLike(t *testing.T) {
	const out = 5.0
	minkowskiDistance := new(MinkowskiDistance)
	distance, err := minkowskiDistance.Distance(v1, v2, 1.0)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if distance != out {
		t.Errorf("Distance is wrong. Wanted %v - got %v", out, distance)
	}
}
func TestMinkowskiDistanceWeighted(t *testing.T) {
	const out = 4.24264
	const precision = 0.00001
	weights := []float64{2., 2., 2.}
	minkowskiDistanceWeighted := new(MinkowskiDistanceWeighted)
	distance, err := minkowskiDistanceWeighted.Distance(v1, v2, weights, 2.0)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestChebyshevDistance(t *testing.T) {
	const out = 2.
	const precision = 0.00001
	chebyshevDistance := new(ChebyshevDistance)
	distance, err := chebyshevDistance.Distance(v1, v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestHammingDistance(t *testing.T) {
	const out = 3.
	const precision = 0.00001
	hammingDistance := new(HammingDistance)
	distance, err := hammingDistance.Distance(v1, v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestBrayCurtisDistance(t *testing.T) {
	const out = 0.384615
	const precision = 0.000001
	brayCurtis := new(BrayCurtisDistance)
	distance, err := brayCurtis.Distance(v1, v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}
func TestCanberraDistance(t *testing.T) {
	const out = 1.166667
	const precision = 0.000001
	canberraDistance := new(CanberraDistance)
	distance, err := canberraDistance.Distance(v1, v2)
	if err != nil {
		t.Errorf("An unexpected error occured %v", err)
	}
	if !isEqual(distance, out, precision) {
		t.Errorf("Distance is wrong. Got %v wanted %v by a precision of %v", distance, out, precision)
	}
}

// http://www.wolframalpha.com/input/?i=52%C2%B0+31%27+0%22+N+13%C2%B0+24%27+0%22+E+distance+35%C2%B0+42%27+0%22+N+139%C2%B0+46%27+0%22+E
// http://www.frustfrei-lernen.de/mathematik/bogenmass-und-gradmass.html
// http://de.wikipedia.org/wiki/Orthodrome#Berechnungsbeispiel_Berlin_.E2.80.93_Tokio
// http://en.wikipedia.org/wiki/Great-circle_distance
//func TestGeoDistance(t *testing.T) {
//
//	const out = 8941
//	input1 := []float64{52.517, 13.40}
//	input2 := []float64{35.70, 139.767}
//	circumstance  := 40000.0 // circumstance of our earth
//	geoDistance   := new(GeoDistance)
//	distance, err := geoDistance.Distance(input1, input2, circumstance)
//
//	if err != nil {
//		t.Errorf("An unexpected error occured %v",err)
//	}
//	if distance != out {
//		t.Errorf("Distance is wrong. Got %v wanted %v", distance, out)
//	}
//}

func TestGeoCoordinateToDecimal(t *testing.T) {
	const input = "52°31'1\"N"
	const output = 52.516944
	const precision = 0.0000001
	d := new(GeoCoordinate)
	d.ParseDMS(input)

	if res := d.ToDecimal(); isEqual(res, output, precision) {
		t.Errorf("Converted %v got %v (%2.6f) wanted %v", input, res, res, output)
	}
}

func TestDecimalToGeo(t *testing.T) {
	output := GeoCoordinate{
		Degree:            52,
		Minutes:           31,
		Seconds:           1.3,
		CardinalDirection: "",
	}
	const input = 52.517028
	d := GeoCoordinate{}
	d.ParseDecimal(input, 2)
	if d != output {
		t.Errorf("Converted %v got %v  wanted %v", input, d, output)
	}
}

func TestParseDMS(t *testing.T) {
	const input_post_n = "52°31'1\"N"
	output_post_n := GeoCoordinate{
		Degree:            52,
		Minutes:           31,
		Seconds:           1,
		CardinalDirection: "N",
	}
	const input_post_e = "13°24'1\"E"
	output_post_e := GeoCoordinate{
		Degree:            13,
		Minutes:           24,
		Seconds:           1,
		CardinalDirection: "E",
	}
	const input_post_s = "14°25'2\"S"
	output_post_s := GeoCoordinate{
		Degree:            14,
		Minutes:           25,
		Seconds:           2,
		CardinalDirection: "S",
	}
	const input_post_w = "14°25'2\"w"
	output_post_w := GeoCoordinate{
		Degree:            14,
		Minutes:           25,
		Seconds:           2,
		CardinalDirection: "W",
	}
	const input_pre_s = "S15°26'3\""
	output_pre_s := GeoCoordinate{
		Degree:            15,
		Minutes:           26,
		Seconds:           3,
		CardinalDirection: "S",
	}
	const input_pre_w = "W16°27'4\""
	output_pre_w := GeoCoordinate{
		Degree:            16,
		Minutes:           27,
		Seconds:           4,
		CardinalDirection: "W",
	}
	const input_pre_n = "n17°28'5\""
	output_pre_n := GeoCoordinate{
		Degree:            17,
		Minutes:           28,
		Seconds:           5,
		CardinalDirection: "N",
	}
	const input_pre_e = "e18°29'6\""
	output_pre_e := GeoCoordinate{
		Degree:            18,
		Minutes:           29,
		Seconds:           6,
		CardinalDirection: "E",
	}

	d := GeoCoordinate{}
	var err = d.ParseDMS(input_post_n)

	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_post_n {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_post_n, d, input_post_n)
	}

	err = d.ParseDMS(input_post_e)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_post_e {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_post_e, d, input_post_e)
	}

	err = d.ParseDMS(input_post_s)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_post_s {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_post_s, d, input_post_s)
	}

	err = d.ParseDMS(input_post_w)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_post_w {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_post_w, d, input_post_w)
	}

	err = d.ParseDMS(input_pre_s)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_pre_s {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_pre_s, d, input_pre_s)
	}

	err = d.ParseDMS(input_pre_w)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_pre_w {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_pre_w, d, input_pre_w)
	}

	err = d.ParseDMS(input_pre_n)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_pre_n {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_pre_n, d, input_pre_n)
	}

	err = d.ParseDMS(input_pre_e)
	if err != nil {
		t.Errorf("Unepected Error %v", err)
	}
	if d != output_pre_e {
		t.Errorf("Convertion faild. Wanted %v got %v for %v", output_pre_e, d, input_pre_e)
	}

}

func TestConvertDMSToDegrees(t *testing.T) {
	//	const inputLati = "52°31'1\"N"
	//	const inputLong = "13° 24' 1\" E"
	//	const outLati = 3.2
	//	const outLong = -64.1
	//	const precision = 0.000001
	//	g := new(GeoDistance)
	//	result := g.ConvertDMSToDegrees(inputLati, inputLong)
	//	if ! isEqual(result[0], outLati, precision) {
	//		t.Errorf("We got for lantitude %v but wanted %v", result[0], outLati)
	//	}
	//	if ! isEqual(result[1], outLong, precision) {
	//		t.Errorf("We got for lantitude %v but wanted %v", result[1], outLong)
	//	}
}

func isEqual(f1, f2 float64, tolerance float64) bool {
	if math.Abs((f1 - f2)) <= tolerance {
		return true
	}
	return false
}
