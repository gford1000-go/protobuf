package encryption

import "errors"

var errDecryptionError = errors.New("error during decryption")

type gcmAlgo struct {
}

// Decrypt attempts to decrypt using the key.
func (g *gcmAlgo) Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
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

// Encrypt attempts to encrypt using the specified key
func (g *gcmAlgo) Encrypt(key []byte, plaintext []byte) ([]byte, AlgoType, error) {
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

// gcmAlgorith implements both Algorithm and AlgorithmCreator
type gcmAlgorithm struct {
	a *gcmAlgo
}

func (g *gcmAlgorithm) GetType() AlgoType {
	return GCM
}

func (g *gcmAlgorithm) GetEncryptor() Encryptor {
	return g.a
}

func (g *gcmAlgorithm) GetDecryptor() Decryptor {
	return g.a
}

func (g *gcmAlgorithm) New() Algorithm {
	return g
}

// NewGCMCreator returns an AlgorithmCreator for GCM
func NewGCMCreator() AlgorithmCreator {
	return &gcmAlgorithm{
		a: &gcmAlgo{},
	}
}
