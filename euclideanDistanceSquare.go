package goDistances

type EuclideanDistanceSquare struct {
}

func (e2d EuclideanDistanceSquare) Distance(params ...interface{}) (float64, error) {
	vector1 := params[0].([]float64)
	vector2 := params[1].([]float64)
	e := new(EuclideanDistance)
	d, err := e.Distance(vector1, vector2)

	return d*d, err
}
