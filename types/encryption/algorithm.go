package encryption

import "errors"

// Algorithm represents available encryption algorithms
type Algorithm uint

const (
	Unknown Algorithm = iota
	GCM
)

var errUnknownAlgorithmUsed = errors.New("unsupported algorithm used for encryption")

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
