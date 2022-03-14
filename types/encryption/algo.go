package encryption

import "errors"

// Algorithm represents available encryption algorithms
type Algorithm uint

const (
	Unknown Algorithm = iota
	GCM
)

var errUnknownAlgorithmUsed = errors.New("unsupported algorithm used for encryption")

// NewAlgo returns the corresponding Algo to the Algorithm,
// or returns Algo_Unknown and an error if not matched
func NewAlgo(a Algorithm) (Algo, error) {
	switch a {
	case GCM:
		{
			return Algo_GCM, nil
		}
	}
	return Algo_UnknownAlgo, errUnknownAlgorithmUsed
}

// ParseAlgo returns the corresponding Algorithm to the Algo,
// or returns Unknown and an error if not matched
func ParseAlgo(a Algo) (Algorithm, error) {
	switch a {
	case Algo_GCM:
		{
			return GCM, nil
		}
	}
	return Unknown, errUnknownAlgorithmUsed
}
