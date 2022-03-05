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

// GetKeys returns the map of key tokens to their associated keys
func (g *GCMTokenKeyEncryptor) GetKeys() map[string][]byte {
	return g.keys
}

// Encrypt attempts to decrypt using the key associated with the token.
func (g *GCMTokenKeyEncryptor) Encrypt(keyToken []byte, plaintext []byte) ([]byte, Algorithm, error) {
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

	gcm, err := NewGCMEncryptor(k)
	if err != nil {
		return nil, Unknown, errEncryptionError
	}

	b, err := gcm.Encrypt(plaintext)
	if err != nil {
		return nil, Unknown, errEncryptionError
	}

	return b, GCM, nil
}
