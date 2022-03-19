package encryption

import "errors"

var errUnknownAlgorithmUsed = errors.New("unsupported algorithm used for encryption")

// NewAlgo returns the corresponding Algo to the AlgoType,
// or returns Algo_Unknown and an error if not matched
func NewAlgo(a AlgoType) (Algo, error) {
	switch a {
	case GCM:
		{
			return Algo_GCM, nil
		}
	}
	return Algo_UnknownAlgo, errUnknownAlgorithmUsed
}

// ParseAlgo returns the corresponding AlgoType to the Algo,
// or returns Unknown and an error if not matched
func ParseAlgo(a Algo) (AlgoType, error) {
	switch a {
	case Algo_GCM:
		{
			return GCM, nil
		}
	}
	return Unknown, errUnknownAlgorithmUsed
}
