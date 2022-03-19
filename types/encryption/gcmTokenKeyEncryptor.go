package encryption

import (
	"errors"
)

var errInvalidKeyToken = errors.New("keyToken not valid")
var errEncryptionError = errors.New("error during encryption")

// NewTokenKeyEncryptor returns a new instance of TokenKeyDecryptor
// corresponding to the specified AlgoType
func NewTokenKeyEncryptor(t AlgoType) (TokenKeyEncryptor, error) {
	a, err := DefaultAlgoFactory.GetAlgorithm(t)
	if err != nil {
		return nil, errUnknownAlgoType
	}
	return &defaultTokenKeyEncryptor{
		a:    a,
		keys: make(map[string][]byte),
	}, nil
}

// defaultTokenKeyEncryptor implements TokenKeyEncryptor for encryption
type defaultTokenKeyEncryptor struct {
	a    Algorithm
	keys map[string][]byte
}

// GetKeys returns the map of key tokens to their associated keys,
// encrypted using the specified key and the specified Algorithm
func (g *defaultTokenKeyEncryptor) GetKeys(key []byte, a Algorithm) (*EncryptedObject, error) {

	k := &Keys{
		Keys: g.keys,
	}

	return NewEncryptedObject(key, k, a.GetEncryptor())
}

// EncryptFromToken attempts to encrypt using the key associated with the token.
func (g *defaultTokenKeyEncryptor) EncryptFromToken(keyToken []byte, plaintext []byte) ([]byte, AlgoType, error) {
	if len(keyToken) == 0 {
		return nil, Unknown, errInvalidKeyToken
	}

	k, ok := g.keys[string(keyToken)]
	if !ok {
		var err error
		k, err = g.a.CreateKey()
		if err != nil {
			return nil, Unknown, err
		}
		g.keys[string(keyToken)] = k
	}

	return g.a.GetEncryptor().Encrypt(k, plaintext)
}
