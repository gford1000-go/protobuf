package encryption

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

var errObjectIsNil = errors.New("EncryptedObject must not be nil")
var errFactoryIsNil = errors.New("AlgorithmFactory must not be nil")
var errAlgoMismatch = errors.New("AlgoType mismatch")
var errMissingKeyToken = errors.New("keyToken not found")

// NewTokenKeyDecryptor returns a new instance of TokenKeyDecryptor,
// prefilled with the specified set of key tokens and associated keys that
// have been retrieved from the encrypted object using the specified key
func NewTokenKeyDecryptor(key []byte, keys *EncryptedObject, factory AlgorithmFactory) (TokenKeyDecryptor, error) {

	if keys == nil {
		return nil, errObjectIsNil
	}

	if factory == nil {
		return nil, errFactoryIsNil
	}

	a, err := ParseAlgo(keys.A)
	if err != nil {
		return nil, err
	}

	algo, err := factory.GetAlgorithm(a)
	if err != nil {
		return nil, err
	}

	b, err := algo.GetDecryptor().Decrypt(key, keys.V)
	if err != nil {
		return nil, err
	}

	k := &Keys{}
	err = proto.Unmarshal(b, k)
	if err != nil {
		return nil, err
	}

	return &gcmTokenKeyDecryptor{
		a:    algo,
		keys: k.GetKeys(),
	}, nil
}

// gcmTokenKeyDecryptor implements TokenKeyDecryptor for GCM symmetric encryption
type gcmTokenKeyDecryptor struct {
	keys map[string][]byte
	a    Algorithm
}

// DecryptFromToken attempts to decrypt using the key associated with the token.
func (g *gcmTokenKeyDecryptor) DecryptFromToken(keyToken []byte, algo AlgoType, ciphertext []byte) ([]byte, error) {
	if algo != g.a.GetType() {
		return nil, errAlgoMismatch
	}

	k, ok := g.keys[string(keyToken)]
	if !ok {
		return nil, errMissingKeyToken
	}

	return g.a.GetDecryptor().Decrypt(k, ciphertext)
}
