package goDistances

type Distancer interface {
	Distance(params ...interface{}) (float64, error)
}
