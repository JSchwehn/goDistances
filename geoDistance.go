package goDistances

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// GeoDistance calculates the distance
type GeoDistance struct{}
type GeoCoordinate struct {
	Degree            float64
	Minutes           float64
	Seconds           float64
	CardinalDirection string // n(orth),(s)outh, (e)ast or (w)est
}

func (gC GeoCoordinate) ToDMS() string {
	return fmt.Sprintf("%s %f°%f'%f\"", gC.CardinalDirection, gC.Degree, gC.Minutes, gC.Seconds)
}
func (gC GeoCoordinate) ToDecimal() float64 {
	decimal := gC.Degree
	decimal += gC.Minutes / 60
	decimal += gC.Seconds / 3600
	return decimal
}
func (gC *GeoCoordinate) ParseDecimal(d float64, round int) error {
	gC.Degree = math.Floor(d)
	leftover := (d - gC.Degree) * 60
	gC.Minutes = math.Floor(leftover)
	gC.Seconds = gC.round((leftover-gC.Minutes)*60, 2)
	return nil
}
func (gC *GeoCoordinate) ParseDMS(dms string) error {
	dmsCoordinate := regexp.MustCompile("^([NSEW])*\\s*(\\d+)\\s*°\\s*(\\d+)\\s*'\\s*(\\d+)\\s*\"\\s*([NSEW])*")
	if match := dmsCoordinate.FindStringSubmatch(strings.ToUpper(dms)); match != nil {
		// detect the cardinal direction char
		if match[5] != "" {
			gC.CardinalDirection = match[5]
		} else {
			gC.CardinalDirection = match[1]
		}
		gC.Degree, _ = strconv.ParseFloat(match[2], 64)
		gC.Minutes, _ = strconv.ParseFloat(match[3], 64)
		gC.Seconds, _ = strconv.ParseFloat(match[4], 64)
		return nil
	}
	return fmt.Errorf("Could not parse given dms coordinates %v", dms)
}
func (gC GeoCoordinate) round(val float64, places int) (newVal float64) {
	var round float64
	roundOn := .5
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// Distance takes two float64 slices which have to have two elements each.
// The first parameter of the slice  will be considered as north angle and the second
// one will be considered as east angle.
// It will return a float for the distance between the two slices.
func (mG GeoDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 3 {
		return -1., fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	d1 := params[0].([]float64) // latitude *longitute
	d2 := params[1].([]float64) //
	c := params[2].(float64)

	fmt.Printf("We got %v %v %v  \n", d1, d2, c)
	d := 0.0

	return d, nil
}
