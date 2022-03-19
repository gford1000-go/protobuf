package encryption

import (
	"errors"
	"sync"
)

var errUnknownAlgorithmUsed = errors.New("unsupported algorithm used for encryption")

var algoMap1 map[AlgoType]Algo
var algoMap2 map[Algo]AlgoType
var algoLock sync.Mutex

func init() {
	algoMap1 = make(map[AlgoType]Algo)
	algoMap2 = make(map[Algo]AlgoType)

	RegisterAlgoMapping(Algo_GCM, GCM)
}

// RgisterAlgoMapping provides the ability to specify new
// mappings between the proto definition and go code
func RegisterAlgoMapping(a Algo, at AlgoType) {
	algoLock.Lock()
	defer algoLock.Unlock()

	algoMap1[at] = a
	algoMap2[a] = at
}

// NewAlgo returns the corresponding Algo to the AlgoType,
// or returns Algo_Unknown and an error if not matched
func NewAlgo(at AlgoType) (Algo, error) {
	algoLock.Lock()
	defer algoLock.Unlock()

	a, ok := algoMap1[at]
	if !ok {
		return Algo_UnknownAlgo, errUnknownAlgorithmUsed
	}
	return a, nil
}

// ParseAlgo returns the corresponding AlgoType to the Algo,
// or returns Unknown and an error if not matched
func ParseAlgo(a Algo) (AlgoType, error) {
	algoLock.Lock()
	defer algoLock.Unlock()

	at, ok := algoMap2[a]
	if !ok {
		return Unknown, errUnknownAlgorithmUsed
	}
	return at, nil
}
