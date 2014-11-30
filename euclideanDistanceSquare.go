package goDistances

// EuclideanDistanceSquare calculates an euclidean distance and square the result
// It will return a float for the distance between the two slices.
type EuclideanDistanceSquare struct {}

// Distance takes two float64 slices which have to have the same size.
func (e2d EuclideanDistanceSquare) Distance(params ...interface{}) (float64, error) {
	vector1 := params[0].([]float64)
	vector2 := params[1].([]float64)
	e := new(EuclideanDistance)
	d, err := e.Distance(vector1, vector2)

	return d*d, err
}
