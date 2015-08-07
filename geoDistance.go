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
	CardinalDirection string // (n)orth,(s)outh, (e)ast or (w)est
}
type GeoPoint struct {
	Latitude  GeoCoordinate
	Longitude GeoCoordinate
}

func (gC *GeoCoordinate) ToDMS() string {
	return fmt.Sprintf("%s %1.0f° %1.0f' %1.2f\"", gC.CardinalDirection, gC.Degree, gC.Minutes, gC.Seconds)
}
func (gC GeoCoordinate) ToDecimal() float64 {
	decimal := gC.Degree
	decimal += gC.Minutes / 60
	decimal += gC.Seconds / 3600

	return decimal
}

func (gC GeoCoordinate) ToRad() float64 {
	deg := gC.ToDecimal()
	return deg * (math.Pi / 180)
}

func (gC *GeoCoordinate) ParseDecimalAsLatitude(d float64) error {

	if err := gC.ParseDecimal(d); err != nil {
		return err
	}
	if d < 0 {
		gC.CardinalDirection = "S"
	} else {
		gC.CardinalDirection = "N"
	}

	return nil
}
func (gC *GeoCoordinate) ParseDecimalAsLongitude(d float64) error {
	if err := gC.ParseDecimal(d); err != nil {
		return err
	}
	if d < 0 {
		gC.CardinalDirection = "W"
	} else {
		gC.CardinalDirection = "E"
	}

	return nil
}

func (gC *GeoCoordinate) ParseDecimal(d float64) error {
	round := 2
	d = math.Abs(d)
	gC.Degree = math.Floor(d)
	leftover := (d - gC.Degree) * 60
	gC.Minutes = math.Floor(leftover)
	gC.Seconds = gC.round((leftover-gC.Minutes)*60, round)

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
	roundOn := .5 //
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

//http://en.wikipedia.org/wiki/Vincenty%27s_formulae
func (mG GeoDistance) Distance(params ...interface{}) (float64, error) {
	if len(params) != 4 {
		return -1., fmt.Errorf("wrong parameter count. Needed 3 got %d", len(params))
	}
	p1 := params[0].(GeoPoint) // latitude *longitute
	p2 := params[1].(GeoPoint) // second point
	r := params[2].(float64)   // radius
	f := params[3].(float64)   // flattening of the ellipsoid

	lat1 := p1.Latitude.ToDecimal()
	long1 := p1.Longitude.ToDecimal()
	lat2 := p2.Latitude.ToDecimal()
	long2 := p2.Longitude.ToDecimal()

	F := (math.Pi / 180) * ((lat1 + lat2) / 2)
	G := (math.Pi / 180) * ((lat1 - lat2) / 2)
	l := (math.Pi / 180) * ((long1 - long2) / 2)

	S := (math.Sin(G)*math.Sin(G))*(math.Cos(l)*math.Cos(l)) +
		(math.Cos(F)*math.Cos(F))*(math.Sin(l)*math.Sin(l))
	C := (math.Cos(G)*math.Cos(G))*(math.Cos(l)*math.Cos(l)) +
		(math.Sin(F)*math.Sin(F))*(math.Sin(l)*math.Sin(l))
	w := math.Atan(math.Sqrt(S / C))
	D := 2 * w * r

	R := (math.Sqrt(S * C)) / w
	H1 := (3*R - 1.0) / (2.0 * C)
	H2 := (3*R + 1.0) / (2.0 * S)
	d := D * (1.0 + f*H1*(math.Sin(F)*math.Sin(F))*(math.Cos(G)*math.Cos(G)) -
		f*H2*(math.Cos(F)*math.Cos(F))*(math.Sin(G)*math.Sin(G)))

	return d, nil
}
