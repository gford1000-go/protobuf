package encryption

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

var errMissingKeyToken = errors.New("keyToken not found")
var errDecryptionError = errors.New("error during decryption")

// NewGCMTokenKeyDecryptor returns a new instance of GCMTokenKeyDecryptor,
// prefilled with the specified set of key tokens and associated keys that
// have been retrieved from the encrypted object using the specified key
func NewGCMTokenKeyDecryptor(key []byte, keys *EncryptedObject) (*GCMTokenKeyDecryptor, error) {

	a, err := ParseAlgo(keys.A)
	if err != nil {
		return nil, err
	}

	g := &GCMTokenKeyDecryptor{}
	b, err := g.Decrypt(key, a, keys.V)
	if err != nil {
		return nil, err
	}

	k := &Keys{}
	err = proto.Unmarshal(b, k)
	if err != nil {
		return nil, err
	}

	return &GCMTokenKeyDecryptor{keys: k.GetKeys()}, nil
}

// GCMTokenKeyDecryptor implements TokenKeyDecryptor for GCM symmetric encryption
type GCMTokenKeyDecryptor struct {
	keys map[string][]byte
}

// DecryptFromToken attempts to decrypt using the key associated with the token.
func (g *GCMTokenKeyDecryptor) DecryptFromToken(keyToken []byte, algo Algorithm, ciphertext []byte) ([]byte, error) {
	k, ok := g.keys[string(keyToken)]
	if !ok {
		return nil, errMissingKeyToken
	}

	return g.Decrypt(k, algo, ciphertext)
}

// Decrypt attempts to decrypt using the key.
func (g *GCMTokenKeyDecryptor) Decrypt(key []byte, algo Algorithm, ciphertext []byte) ([]byte, error) {
	if algo != GCM {
		return nil, errDecryptionError
	}

	gcm, err := NewGCMEncryptor(key)
	if err != nil {
		return nil, errDecryptionError
	}

	b, err := gcm.Decrypt(ciphertext)
	if err != nil {
		return nil, errDecryptionError
	}

	return b, nil
}
