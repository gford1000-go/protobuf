package encryption

import (
	"errors"
)

var errInvalidKeyToken = errors.New("keyToken not valid")
var errEncryptionError = errors.New("error during encryption")

// NewGCMTokenKeyEncryptor returns a new instance of GCMTokenKeyDecryptor
func NewGCMTokenKeyEncryptor() *GCMTokenKeyEncryptor {
	return &GCMTokenKeyEncryptor{keys: make(map[string][]byte)}
}

// GCMTokenKeyEncryptor implements TokenKeyEncryptor for GCM symmetric encryption
type GCMTokenKeyEncryptor struct {
	keys map[string][]byte
}

// GetKeys returns the map of key tokens to their associated keys,
// encrypted using the specified key and the GCM algorithm
func (g *GCMTokenKeyEncryptor) GetKeys(key []byte) (*EncryptedObject, error) {

	k := &Keys{
		Keys: g.keys,
	}

	return NewEncryptedObject(key, k, g)
}

// EncryptFromToken attempts to encrypt using the key associated with the token.
func (g *GCMTokenKeyEncryptor) EncryptFromToken(keyToken []byte, plaintext []byte) ([]byte, AlgoType, error) {
	if len(keyToken) == 0 {
		return nil, Unknown, errInvalidKeyToken
	}

	k, ok := g.keys[string(keyToken)]
	if !ok {
		var err error
		k, err = NewAESKey()
		if err != nil {
			return nil, Unknown, err
		}
		g.keys[string(keyToken)] = k
	}

	return g.Encrypt(k, plaintext)
}

// Encrypt attempts to encrypt using the specified key
func (g *GCMTokenKeyEncryptor) Encrypt(key []byte, plaintext []byte) ([]byte, AlgoType, error) {
	gcm, err := NewGCMEncryptor(key)
	if err != nil {
		return nil, Unknown, errEncryptionError
	}

	b, err := gcm.Encrypt(plaintext)
	if err != nil {
		return nil, Unknown, errEncryptionError
	}

	return b, GCM, nil
}
