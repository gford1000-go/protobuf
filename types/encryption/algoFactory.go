package encryption

import (
	"errors"
)

var errInvalidAlgoCreator = errors.New("AlgorihmCreator must not be nil")
var errInvalidAlgoType = errors.New("AlgorithmCreator has invalid AlgoType")
var errUnknownAlgoType = errors.New("unknown AlgoType requested")

func init() {
	DefaultAlgoFactory, _ = NewAlgorithmFactory([]AlgorithmCreator{
		NewGCMCreator(),
	})
}

// DefaultAlgoFactory is a AlgoFactory pre-filled with existing AlgoTypes,
// currently GCM
var DefaultAlgoFactory *AlgoFactory

// NewAlgorithmFactory returns an instance of AlgoFactory, pre-filled with
// the specified set of AlgorithmCreators
func NewAlgorithmFactory(as []AlgorithmCreator) (*AlgoFactory, error) {
	f := &AlgoFactory{
		m: make(map[AlgoType]AlgorithmCreator),
	}

	for _, c := range as {
		if err := f.AddAlgorithmCreator(c); err != nil {
			return nil, err
		}
	}

	return f, nil
}

// AlgoFactory manufactures instances of Algorithm by invoking the
// AlgorithmCreator for the required AlgoType
type AlgoFactory struct {
	m map[AlgoType]AlgorithmCreator
}

// AddAlgorithmCreator inserts the specified AlgorithmCreator
// into the AlgoFactory
func (f *AlgoFactory) AddAlgorithmCreator(c AlgorithmCreator) error {
	if c == nil {
		return errInvalidAlgoCreator
	}

	a := c.New()
	if a == nil {
		return errInvalidAlgoCreator
	}

	if a.GetType() == "" {
		return errInvalidAlgoType
	}

	// Will replace existing for the same AlgoType
	f.m[a.GetType()] = c

	return nil
}

// GetAlgorithm returns an instance of a Algorithm of the
// specified AlgoType
func (f *AlgoFactory) GetAlgorithm(t AlgoType) (Algorithm, error) {
	c, ok := f.m[t]
	if !ok {
		return nil, errUnknownAlgoType
	}
	return c.New(), nil
}
