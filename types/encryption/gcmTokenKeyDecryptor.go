package encryption

import (
	"errors"
)

var errMissingKeyToken = errors.New("keyToken not found")
var errDecryptionError = errors.New("error during decryption")

// NewGCMTokenKeyDecryptor returns a new instance of GCMTokenKeyDecryptor,
// prefilled with the specified set of key tokens and associated keys
func NewGCMTokenKeyDecryptor(keys map[string][]byte) *GCMTokenKeyDecryptor {
	return &GCMTokenKeyDecryptor{keys: keys}
}

// GCMTokenKeyDecryptor implements TokenKeyDecryptor for GCM symmetric encryption
type GCMTokenKeyDecryptor struct {
	keys map[string][]byte
}

// DecryptFromToken attempts to decrypt using the key associated with the token.
func (g *GCMTokenKeyDecryptor) DecryptFromToken(keyToken []byte, algo Algorithm, ciphertext []byte) ([]byte, error) {
	if algo != GCM {
		return nil, errDecryptionError
	}

	k, ok := g.keys[string(keyToken)]
	if !ok {
		return nil, errMissingKeyToken
	}

	gcm, err := NewGCMEncryptor(k)
	if err != nil {
		return nil, errDecryptionError
	}

	b, err := gcm.Decrypt(ciphertext)
	if err != nil {
		return nil, errDecryptionError
	}

	return b, nil
}
