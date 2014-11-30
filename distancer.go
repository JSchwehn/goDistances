package goDistances

// Distancer Interface with a takes n-parameters, the implementation has to take care that the
// parameters are valid.
type Distancer interface {
	Distance(params ...interface{}) (float64, error)
}
